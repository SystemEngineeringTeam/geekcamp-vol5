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
              ><v-card-title>お疲れ様です!</v-card-title>
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
                  href="http://twitter.com/share?url=https://localhost:8080/&text=瞑想によってドーパミンが放出されました！！！&hashtags=ドーパ民"
                >
                  <v-icon
                    dark
                    size="42px"
                    style="text-decoration: none"
                    color="white"
                    >mdi-twitter</v-icon
                  >ツイート</v-btn
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
