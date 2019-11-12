var store = new Vuex.Store({

    state: {
      connected: null,
      logged: null,
      service: null,
      version: null,
      history: null,
      currentNode: null,
      currentEdge: null,
      emphasizedNodes: [],
      highlightedNodes: [],
      highlightInprogress: new Map(),
      notifications: [],
      topologyFilter: "",
      topologyHighlight: "",
      topologyTimeContext: 0,
      token:""
    },
  
    mutations: {
      history: function(state, support) {
        state.history = support;
      },
  
      topologyFilter: function(state, filter) {
        state.topologyFilter = filter;
      },
  
      topologyHighlight: function(state, filter) {
        state.topologyHighlight = filter;
      },
  
      topologyTimeContext: function(state, time) {
        state.topologyTimeContext = time;
      },
  
      login: function(state, data) {
        state.token = data;
        state.logged = true;
        state.permissions = "";
      },
  
      logout: function(state) {
        state.logged = false;
        state.token = "";
        state.permissions = [];
  
      },
  
      connected: function(state) {
        state.connected = true;
      },
  
      disconnected: function(state) {
        state.connected = false;
      },
  
      version: function(state, version) {
        state.version = version;
      },
  
      addNotification: function(state, notification) {
        if (state.notifications.length > 0 &&
            state.notifications.some(function(n) {
              return n.message === notification.message;
            })) {
          return;
        }
        state.notifications.push(notification);
      },
  
      removeNotification: function(state, notification) {
        state.notifications = state.notifications.filter(function(n) {
          return n !== notification;
        });
      },
  
    },
  
  });

  
var routes = [
    { path: '/', component: MainComponent  },
    { path: '/home', redirect : '/'}
  ];

var router = new VueRouter({
    mode: 'history',
    linkActiveClass: 'active',
    routes: routes
});

var app = new Vue({
    router: router,
  
    store: store,
  
    mixins: [notificationMixin],
  
    created: function() {
      var self = this;
  
      //this.setThemeFromConfig();
  
      //this.checkAPI();
  
      this.interval = null;
  
      // global handler to detect authorization errors
      /*
      $(document).ajaxError(function(evt, e) {
        switch (e.status) {
          case 401:
            self.$error({message: 'Authentication failed'});
            self.$store.commit('logout');
            break;
        }

  
        return e;
      });
      */
    },
  });

$(document).ready(function() {
    Vue.config.devtools = true;  
    app.$mount('#app');
//    app.setThemeFromConfig();
  });