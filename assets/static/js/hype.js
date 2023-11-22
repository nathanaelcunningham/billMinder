htmx.defineExtension("slide", {
  onEvent: function(name, evt) {
      console.log(name)
    if (name === "htmx:configRequest") {
      if (evt.detail.elt.hasAttribute("hx-slide-in")) {
        evt.detail.elt.style.transform = "translateX(0)";
      }
      if (evt.detail.elt.hasAttribute("hx-slide-out")) {
        evt.detail.elt.style.transform = "translateX(100%)";
      }
    }
  },
});

