<template>
  <body>
    <center class="all">
      <p color="black" class="texttmp">何分間瞑想しますか？</p>
      <v-row class="pt-10">
        <v-col>
          <v-btn
            large
            elevation="4"
            height="5.5rem"
            width="12rem"
            style="text-transform: none"
            @click="NowTimeFunction(2)"
          >
            <div class="font">2min</div>
          </v-btn>
        </v-col>
        <v-col>
          <v-btn
            large
            elevation="4"
            height="5.5rem"
            width="12rem"
            style="text-transform: none"
            @click="NowTimeFunction(4)"
          >
            <div class="font">4min</div>
          </v-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn
            large
            elevation="4"
            height="5.5rem"
            width="12rem"
            style="text-transform: none"
            @click="NowTimeFunction(6)"
          >
            <div class="font">6min</div>
          </v-btn>
        </v-col>
        <v-col>
          <v-btn
            large
            elevation="4"
            height="5.5rem"
            width="12rem"
            style="text-transform: none"
            @click="NowTimeFunction(8)"
          >
            <div class="font">8min</div>
          </v-btn>
        </v-col>
      </v-row>
    </center>
  </body>
</template>

<script>
export default {
  data() {
    return {
      min2: 2,
      min4: 4,
      min6: 6,
      min8: 8,
    }
  },
  methods: {
    NowTimeFunction(time) {
      switch (time) {
        case 2:
          this.$store.commit('saveMidTimeTmp', this.min2)
          break
        case 4:
          this.$store.commit('saveMidTimeTmp', this.min4)
          break
        case 6:
          this.$store.commit('saveMidTimeTmp', this.min6)
          break
        case 8:
          this.$store.commit('saveMidTimeTmp', this.min8)
          break
        default:
          break
      }
      this.$router.push('/route/Medi2')
    },
  },
}
</script>

<style>
.texttmp {
  color: black;
}
.all {
  margin-top: 10px;
  font-size: 6vh;
}
.font {
  font-size: 32px;
  font-weight: bold;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
</style>