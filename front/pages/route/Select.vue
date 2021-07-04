<template>
  <div>
    <v-container justify="center" style="positon: relative">
      <v-row justify="center" style="positon: relative" align-content="center">
        <v-col cols="12" md="12" class="mt-30">
          <!-- class black-textで文字の色を変更している -->
          <v-btn
            block
            color="orange"
            height="120px"
            class="black--text text-h4"
            rounded
            nuxt
            to="/route/Run"
            >走る</v-btn
          >
        </v-col>
        <v-col cols="12" class="select">
          <v-btn
            block
            color="blue"
            height="120px"
            class="black--text text-h4"
            rounded
            nuxt
            to="/route/Music"
            >音楽を聞く</v-btn
          >
        </v-col>
        <v-col cols="12" class="select">
          <v-btn
            block
            height="120px"
            class="black--text text-h4 test"
            rounded
            nuxt
            to="/route/Meditation"
            >瞑想する</v-btn
          >
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  setup() {},
}
</script>

<style scoped>
.whole {
  position: absolute;
  top: 30px;
}
.select {
  margin: 30px;
}
.test {
  margin: 50px;
  padding: 10px;
  width: 50px;
  border-radius: 5px;
  box-shadow: 5px 5px 10px #bbbbbb, -5px -5px 10px #ffffff;
}
</style>
