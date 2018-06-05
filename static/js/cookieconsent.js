function eraseCookie(name) {
    document.cookie = name+'=; Max-Age=-99999999;';
}
function eraseCookies() {
    eraseCookie('_ga');
    eraseCookie('_gid');
    eraseCookie('_gat_gtag_UA_113732834_1');
}

window.addEventListener("load", function(){
    window.cookieconsent.initialise({
        "palette": {
            "popup": {
                "background": "#000"
            },
            "button": {
                "background": "#fff"
            }
        },
        "position": "bottom-left",
        "type": "opt-in",
        "content": {
            "href": "https://thngstruction.online/privacy.html"
        },
        onInitialise: function (status) {
            var type = this.options.type;
            var didConsent = this.hasConsented();
            if (type == 'opt-out' && !didConsent) {
                eraseCookies();
            }
        },
        onStatusChange: function(status, chosenBefore) {
            var type = this.options.type;
            var didConsent = this.hasConsented();
            if (type == 'opt-out' && !didConsent) {
                eraseCookies();
            }
        },

        onRevokeChoice: function() {
            var type = this.options.type;
            if (type == 'opt-out') {
                eraseCookies();
            }
        }
    })});
