<template>
  <body>
    <center class="all pt-10 mt-10">
      <p>ガイドに従い</p>
      <p>瞑想を行ってください...</p>
      <v-card class="mx-auto" max-width="600">
        <v-card-text>
          <v-row class="mb-4" justify="space-between">
            <v-col class="text-center">
              <span class="text-h1 font-weight-light" v-text="bpm"></span>
              <span class="subheading font-weight-light mr-1">S</span>
            </v-col>
          </v-row>
          <v-slider
            v-model="bpm"
            :color="color"
            track-color="grey"
            always-dirty
            min="0"
            :max="allseconds"
          >
            <audio
              controls
              autoplay
              controlslist="nodownload"
              src="@/assets/healing3.mp3"
            ></audio>
          </v-slider>
        </v-card-text>
      </v-card>
    </center>
    <div>
      <Celebration v-if="judge"></Celebration>
    </div>
  </body>
</template>

<script>
import sound from '~/assets/healing1.mp3'
export default {
  data: () => ({
    bpm: 0,
    maxbpm: 0,
    alltime: 0,
    allseconds: 0,
    judge: false,
  }),

  computed: {
    color() {
      if (this.bpm < this.allseconds / 4) return 'indigo'
      if (this.bpm < (this.allseconds / 4) * 2) return 'teal'
      if (this.bpm < (this.allseconds / 4) * 2.5) return 'green'
      if (this.bpm < (this.allseconds / 4) * 3) return 'orange'
      return 'red'
    },
  },
  created() {
    setInterval(this.coutUp, 1000)

    this.alltime = this.$store.state.MidTimeTmp
    switch (this.alltime) {
      case 2:
      case 4:
      case 6:
      case 8:
        this.allseconds = this.alltime * 60
        break
      default:
        break
    }
    const audio = new Audio(sound)
    audio.play()
  },
  methods: {
    coutUp() {
      this.bpm += 1
      if (this.bpm === this.allseconds) {
        this.judge = true
      }
    },
  },
}
</script>

<style>
.all {
  color: black;
  font-size: 32px;
  font-weight: bold;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
</style>
