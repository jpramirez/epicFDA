/* jshint multistr: true */

var MainComponent = {
    name :"MainComponent",
    mixins: [notificationMixin],
    template: '\
    <div class="container"> \
        <NavigationComponent></NavigationComponent>\
        <br> \
        <br> \
        <div id="why"> \
        </div>\
        <br> \
        <hr> \
        <br> \
        <UploadFormComponent></UploadFormComponent> \
        <br> \
        <br> \
        <ProductOfferingComponent></ProductOfferingComponent> \
        <FooterComponent></FooterComponent> \
    </div> \
    '
  }
  