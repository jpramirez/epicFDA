var app = new Vue({
	delimiters: ['${', '}'],
	el: '#contact',
	data: {
			returnStatus :"",
			returnURL :"",
			uploadPercentage: 0,
			showprogress: false
	},
	methods: {
		submitForm: function () {
			var self = this;

            /*
                Add the form data we need to submit
            */
			var formData = new FormData();
			/*
			Iterate over any file sent over appending the files
			to the form data.
			*/
			var fileForm  = document.getElementById('files');
			var name  = document.getElementById('name').value;
			var email  = document.getElementById('email').value;

           var metadata = {
                'name':  name,
                'email': email
            };

            formData.append('metadata', new Blob([JSON.stringify(metadata)], {type: 'application/json'}));
			for( var i = 0; i < fileForm.files.length; i++ ){
					let file = fileForm.files[i];
					formData.append('files[' + i + ']', file);
			}
			let url = '/api/v1/process/split/upload';
			this.showprogress = true;
			axios.post( url,
				formData,
				{
					     onUploadProgress: function( progressEvent ) {
        				 this.uploadPercentage = parseInt( Math.round( ( progressEvent.loaded * 100 ) / progressEvent.total ) );
     			 }.bind(this)
				}
				).then(function(data, err){
					console.log(data)
					self.data = data;
					self.returnStatus = self.data.data.FileStatus[0].Status
					if (self.returnStatus != "Rejected") {
						self.returnURL = "/api/v1/file/" +self.data.data.FileStatus[0].Hash;
					}
	
				})
				.catch(function(err){
						self.returnStatus = "An error ocurred. Contact the owner :) "
				});

	}
	}
})
