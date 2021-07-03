<template>
  <div class="text-center">
    <v-row>
      <v-col cols="12" md="7">
        <v-card max-width="2000px" class="ma-12 rounded-card">
          <v-row justify="center">
            <v-card-title>{{ target }}m走る</v-card-title>
          </v-row>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-progress-circular
          :rotate="270"
          :size="320"
          :width="30"
          :value="value"
          color="teal"
          button
        >
          <v-row>
            <v-col cols="12">
              <div class="distance">{{ value }}</div>
              <v-btn @click="getLocation()">距離を測る</v-btn>
            </v-col>
            <v-col cols="12">
              <v-btn @click="resetDistance()">リセットする</v-btn>
              {{ testTotal }}km
            </v-col>
          </v-row>
        </v-progress-circular>
      </v-col>
      <v-col cols="12" style="position: relative">
        <v-btn fab class="button-map" nuxt to="Map">
          <v-icon>mdi-map-marker</v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      interval: {},
      value: 0,
      latitude: 35,
      longitude: 138,
      myLatLng1: { lat: 35.397, lng: 138.644 },
      target: this.$store.getters.getTarget,
      testTotal: this.$store.getters.getTotalLength,
    }
  },
  computed: {
    len() {
      return this.$store.getters.getTotalLength
    },
    targetLen() {
      return this.$store.getters.getTarget
    },
  },
  watch: {
    len() {
      console.log('変更されました')
      this.testTotal = this.$store.getters.getTotalLength
    },
    targetlen() {
      this.target = this.$store.getters.getTarget
    },
  },
  beforeDestroy() {
    clearInterval(this.interval)
  },
  mounted() {
    if (this.target === 0) {
      this.randamTarget()
    }
    this.interval = setInterval(() => {
      if (this.value === 100) {
        return (this.value = 0)
      }
      this.value += 1
    }, 20000)
    this.interval = setInterval(() => {
      this.getLocation()
    }, 20000)
  },
  methods: {
    getLocation() {
      navigator.geolocation.getCurrentPosition(this.success, this.error)
    },
    success(position) {
      // 現在の緯度経度を代入
      // console.log(this.$store.getters.getPastLatLng)

      this.myLatLng1.lat = position.coords.latitude
      this.myLatLng1.lng = position.coords.longitude

      position = { lat: this.myLatLng1.lat, lng: this.myLatLng1.lng }

      const pastLatLng = this.$store.getters.getPastLatLng
      console.log(pastLatLng)
      // if (this.$store.getters.getTotalLength === 0) {
      //   // 初めに取得する値は0にするためその場の位置を代入
      //   pastLatLng = { lat: this.myLatLng1.lat, lng: this.myLatLng1.log }
      // } else {
      //   // 一つ前の緯度経度を代入
      //   pastLatLng = this.$store.getters.getPastLatLng
      // }

      // 距離
      const totalLength =
        this.$store.getters.getTotalLength +
        this.distanced(
          this.myLatLng1.lat,
          this.myLatLng1.lng,
          pastLatLng.lat,
          pastLatLng.lng
        )
      console.log(totalLength)
      // 合計距離をローカルストレージに保存
      this.$store.commit('saveTotalLength', totalLength)
      // 前回のデータが上書きされるのでここまでに距離の計算をする

      // 現在の緯度経度の保存
      this.$store.commit('savePosition', position)
    },
    error(err) {
      console.log(err)
    },
    distanced(lat1, lng1, lat2, lng2) {
      lat1 *= Math.PI / 180
      lng1 *= Math.PI / 180
      lat2 *= Math.PI / 180
      lng2 *= Math.PI / 180

      return (
        6371 *
        Math.acos(
          Math.cos(lat1) * Math.cos(lat2) * Math.cos(lng2 - lng1) +
            Math.sin(lat1) * Math.sin(lat2)
        )
      )
    },
    resetDistance() {
      localStorage.clear()
      this.testTotal = 0
    },
    randamTarget() {
      const len = (Math.floor(Math.random() * (10 + 1 - 1)) + 1) * 100
      console.log(len)
      this.$store.commit('setTarget', len)
    },
  },
}
</script>

<style scoped>
.v-progress-circular {
  margin: 1rem;
}
.distance {
  font-size: 40px;
}
.targetRun {
  margin: 30px;
}
.rounded-card {
  border-radius: 50px;
}
.button-map {
  position: absolute;
  right: 20px;
}
</style>
