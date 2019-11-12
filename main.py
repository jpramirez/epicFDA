from bs4 import BeautifulSoup
import requests

# point to output directory
outpath = 'FDA_DATA/'
url = 'https://open.fda.gov/downloads/'
mbyte=1024*1024

print ('Reading: ', url)
html = requests.get(url).text
print (html)
soup = BeautifulSoup(html)

print ('Processing: ', url)
for name in soup.findAll('a', href=True):
    zipurl = name['href']
    if( zipurl.endswith('.zip') ):
        outfname = outpath + zipurl.split('/')[-1]
        r = requests.get(zipurl, stream=True)
        if( r.status_code == requests.codes.ok ) :
            fsize = int(r.headers['content-length'])
            print ('Downloading %s (%sMb)' % ( outfname, fsize/mbyte ))
            with open(outfname, 'wb') as fd:
                for chunk in r.iter_content(chunk_size=1024): # chuck size can be larger
                    if chunk: # ignore keep-alive requests
                        fd.write(chunk)
                fd.close()