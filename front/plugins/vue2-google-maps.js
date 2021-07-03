import Vue from 'vue'
import * as VueGoogleMaps from 'vue2-google-maps'

Vue.use(VueGoogleMaps, {
  load: {
    key: 'AIzaSyCo0sZxINnYofAmIrulMaJ5ncwmyiFLSbI',
    libraries: 'places',
    language: 'ja',
    reginon: 'JP',
    // OR: libraries: 'places,drawing'
    // OR: libraries: 'places,drawing,visualization'
    // (as you require)

    // v: '3.26',
  },

  installComponents: true,
})
