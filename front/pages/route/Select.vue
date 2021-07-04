<template>
  <div>
    <v-container justify="center" style="positon: relative">
      <v-row justify="center" style="positon: relative" align-content="center">
        <p class="mt-5">Please Select</p>
          <v-col v-for="text in texts" :key="text" cols="12" md="12" class="mt-10">
            <NuxtLink to="/route/Meditation">
            <button
              :style="text.color"
              >{{ text.text }}</button
            >
            </NuxtLink>
          </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  data(){
    return {
      texts: [
        {
          text: '走る',
          color: 'color:rgba(247, 70, 70, 0.8)'
        },
        {
          text: '音楽を聞く',
          color: ''
        },
        {
          text:'瞑想する',
          color: 'color:rgba(25, 197, 10, 0.918)'
        }
      ]
    }
  },
  setup() {},
}
</script>

<style scoped>
* {
  text-decoration: none;
}
.whole {
  position: absolute;
  top: 30px;
}
.select {
  margin: 30px;
}
.a{
  color: rgba(25, 197, 10, 0.918)
}
p {
  font-family: 'Roboto', sans-serif;
  text-align: left !important;
  color: #747474;
  text-shadow: 4px 4px 3px #bebebe, -3px -3px 5px #ffffff;
  font-size: 35px;

  text-indent: 0.2em;
  font-weight: 800;
}
button {
  width: 320px;
  height: 100px;
  padding: 1em 2em;
  background: #efefef;
  border: none;
  border-radius: .5rem;
  font-size: 2rem;
  font-weight: 700;
  letter-spacing: .2rem;
  text-align: center;
  outline: none;
  cursor: pointer;
  transition: .2s ease-in-out;
  box-shadow: -6px -6px 14px rgba(255, 255, 255, .7),
              -6px -6px 10px rgba(255, 255, 255, .5),
              6px 6px 8px rgba(255, 255, 255, .075),
              6px 6px 10px rgba(0, 0, 0, .15);
}
button:hover {
  box-shadow: -2px -2px 6px rgba(255, 255, 255, .6),
              -2px -2px 4px rgba(255, 255, 255, .4),
              2px 2px 2px rgba(255, 255, 255, .05),
              2px 2px 4px rgba(0, 0, 0, .1);
}
button:active {
  box-shadow: inset -2px -2px 6px rgba(255, 255, 255, .7),
              inset -2px -2px 4px rgba(255, 255, 255, .5),
              inset 2px 2px 2px rgba(255, 255, 255, .075),
              inset 2px 2px 4px rgba(0, 0, 0, .15);
}
/* .segment {
  padding: 16px*2 0;
  text-align: center;
} */
</style>
