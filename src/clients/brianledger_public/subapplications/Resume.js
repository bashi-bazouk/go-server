// Generated by CoffeeScript 1.10.0
(function() {
  var SubApplication,
    bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; },
    extend = function(child, parent) { for (var key in parent) { if (hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; },
    hasProp = {}.hasOwnProperty;

  SubApplication = require("../framework/SubApplication.coffee").SubApplication;

  exports.Resume = (function(superClass) {
    extend(Resume, superClass);

    function Resume(pure) {
      this.pure = pure != null ? pure : $("<div id=\"resume\" style=\"max-width: 0px; max-height: 0px; padding: 0px; margin: 0px; border: none\">\n	<iframe src=\"/cdn/documents/Brian P Ledger Resume.pdf\"></iframe>\n</div>\n");
      this.is_closed = bind(this.is_closed, this);
      this.close = bind(this.close, this);
      this.open = bind(this.open, this);
    }

    Resume.prototype.open = function(callback) {
      if (this.pure.css('width') !== "0px") {
        return typeof callback === "function" ? callback() : void 0;
      }
      return this.pure.animate({
        'max-width': "21.6cm",
        'max-height': "30.5cm",
        margin: "1px",
        padding: "4px"
      }, 300, (function(_this) {
        return function() {
          _this.pure.css({
            border: "1px solid gray"
          });
          return typeof callback === "function" ? callback() : void 0;
        };
      })(this));
    };

    Resume.prototype.close = function(callback) {
      if (this.pure.css('width') === "0px") {
        return typeof callback === "function" ? callback() : void 0;
      }
      return this.pure.animate({
        'max-width': 0,
        'max-height': 0,
        margin: 0,
        padding: 0
      }, 300, (function(_this) {
        return function() {
          _this.pure.css({
            border: "none"
          });
          return typeof callback === "function" ? callback() : void 0;
        };
      })(this));
    };

    Resume.prototype.is_closed = function() {
      return this.pure.css('width') === "0px";
    };

    return Resume;

  })(SubApplication);

}).call(this);

//# sourceMappingURL=Resume.js.map