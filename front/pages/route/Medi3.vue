<template>
  <div>
    <center class="pt-10 mt-10 text">
      <p class="toptext">ドーパミンが生成されました！！！</p>
      <v-btn
        large
        elevation="1"
        color="#1DA1F2"
        height="3.2rem"
        style="text-transform: none"
      >
        <div style="text-decoration: none">
          <a
            href="http://twitter.com/share?url=https://localhost:8080/&text=瞑想によってドーパミンが放出されました！！！&hashtags=ドーパ民"
            target="_blank"
            style="text-decoration: none; color: white"
          >
            <v-icon dark size="42px" style="text-decoration: none" color="white"
              >mdi-twitter</v-icon
            >
            Twitter
          </a>
        </div>
      </v-btn>
    </center>
  </div>
</template>

<script>
export default {
  data: () => ({
    time: 0,
  }),

  created() {
    setInterval(this.coutUp, 1000)
  },
  methods: {
    coutUp() {
      this.time += 1
      if (this.time == 15) {
        this.$router.push('/')
      }
    },
  },
}
</script>

<style>
.text {
  color: black;
}
.toptext {
  font-size: 70px;
}
</style>