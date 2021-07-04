<template>
  <v-row justify="center">
    <v-col cols="12">
      <!-- <v-btn class="white--text" color="teal" @click="overlay = !overlay">
        Show Overlay
      </v-btn> -->
      <v-overlay :z-index="zIndex" :value="overlay">
        <v-card height="150px">
          <v-row>
            <v-col
              ><v-card-title>目標達成おめでとう!</v-card-title>
              <v-card-actions>
                <v-btn
                  rounded
                  nuxt
                  to="/"
                  class="white--text"
                  color="accent"
                  @click="
                    overlay = false
                    resetDistance()
                  "
                >
                  タイトルへ
                </v-btn>
                <v-btn
                  rounded
                  class="white--text"
                  color="primary"
                  @click="overlay = false"
                  >ツイートする</v-btn
                >
              </v-card-actions></v-col
            >
          </v-row>
        </v-card>
      </v-overlay>
    </v-col>
  </v-row>
</template>

<script>
export default {
  data: () => ({
    overlay: true,
    zIndex: 0,
  }),
  methods: {
    resetDistance() {
      localStorage.clear()
      this.$store.commit('clear')
    },
  },
}
</script>

<style scoped>
.rounded-card {
  border-radius: 100px;
}
</style>
