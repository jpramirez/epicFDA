package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	constants "github.com/jpramirez/epicFDA/pkg/constants"
	fetcher "github.com/jpramirez/epicFDA/pkg/fetcher"
	models "github.com/jpramirez/epicFDA/pkg/models"
	utils "github.com/jpramirez/epicFDA/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "epicFetcher",
	Short: "Fetcher for EPIC FDA API",
	Long:  `A Fast and Flexible on FDA Analysis Platform`,
	RunE:  runFetcher,
}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var config models.Config
var cfgFile string
var projectBase string
var monitorMode string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/config.json)")
	rootCmd.PersistentFlags().Bool("default", true, "Use default configuration")
}

func runFetcher(cmd *cobra.Command, args []string) error {
	config, err := utils.LoadConfiguration(cfgFile)
	if err != nil {
		log.Fatalln("Error and Exiting")
	}
	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	fmt.Println(config.LogFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	_fetch, err := fetcher.NewFetcher(config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()

	_fetch.DownloadAnimalAndVeterinary(results.Results.AnimalAndVeterinary)
	results.SaveIndex(config.DataSetFolder)
	
	_fetch.DownloadFood(results.Results.Food)
	_fetch.DownloadDrug(results.Results.Drug)
	_fetch.DownloadDevice(results.Results.Device)

	fmt.Println("Finished")

	return err
}
