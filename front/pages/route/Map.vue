<template>
  <div>
    <h1>Google Map</h1>
    <div ref="map" style="height: 500px; width: 800px"></div>
    <v-btn @click="getLocation()">現在の位置</v-btn>
  </div>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      map: null,
      marker: null,
      latitude: 35,
      longitude: 138,
      myLatLng: { lat: 35.397, lng: 138.644 },
    }
  },
  head: {
    script: [
      {
        src: 'https://maps.google.com/maps/api/js?key=AIzaSyCo0sZxINnYofAmIrulMaJ5ncwmyiFLSbI&callback=getLocation&libraries=places&language=ja',
      },
    ],
  },
  created() {},
  mounted() {
    const timer = setInterval(() => {
      if (window.google) {
        clearInterval(timer)
        const latlng = new google.maps.LatLng(35.539001, 134.228468)
        this.map = new window.google.maps.Map(this.$refs.map, {
          center: latlng,
          zoom: 8,
          gestureHandling: 'greedy',
        })
        const mopts = {
          position: latlng,
          map: this.map,
          title: 'home',
        }
        const marker = new google.maps.Marker(mopts)

        console.log(this.map)
        console.log(this.marker)
      }
    }, 500)
  },
  methods: {
    getLocation() {
      navigator.geolocation.getCurrentPosition(this.success, this.error)
    },
    success(position) {
      this.myLatLng.lat = position.coords.latitude
      this.myLatLng.lng = position.coords.longitude
      const timer = setInterval(() => {
        if (window.google) {
          clearInterval(timer)
          const opts2 = {
            zoom: 16,
            gestureHandling: 'greedy',
          }

          // ズームを変更
          this.map.setOptions(opts2)

          // 現在の位置の指定
          const latlng = new google.maps.LatLng(
            this.myLatLng.lat,
            this.myLatLng.lng
          )
          // アニメーションによる遷移
          this.map.panTo(latlng)

          // マーカー再セット
          const mopts1 = {
            position: latlng,
            map: this.map,
          }

          const marker = new google.maps.Marker(mopts1)
        }
      }, 500)
    },
    error(err) {
      console.log(err)
    },
  },
}
</script>

<style scoped>
.map {
  width: 100%;
  height: 50%;
}
</style>
