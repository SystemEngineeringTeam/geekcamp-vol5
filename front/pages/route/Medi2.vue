<template>
  <body>
    <center class="all">
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
            max="218"
          >
          </v-slider>
        </v-card-text>
      </v-card>
    </center>
  </body>
</template>
  
<script>
export default {
  data: () => ({
    bpm: 0,
    alltime: 0,
  }),

  computed: {
    color() {
      if (this.bpm < 100) return 'indigo'
      if (this.bpm < 125) return 'teal'
      if (this.bpm < 140) return 'green'
      if (this.bpm < 175) return 'orange'
      return 'red'
    },
    animationDuration() {
      return `${60 / this.bpm}s`
    },
  },
  created() {
    this.alltime = this.$store.state.MidTimeTmp
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