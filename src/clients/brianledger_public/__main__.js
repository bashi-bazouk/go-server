// Generated by CoffeeScript 1.10.0
(function() {
  var Application;

  require("./__init__.coffee");

  Application = require("./Application.coffee").Application;

  $(document).ready(function() {
    var application;
    application = new Application();
    return $("body").append(application.pure);
  });

}).call(this);

//# sourceMappingURL=__main__.js.map