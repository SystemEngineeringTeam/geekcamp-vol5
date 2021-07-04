<template>
  <v-row justify="center" align="center">
    <v-col cols="12">
      <h1>Dopamine App</h1>
    </v-col>
    <v-col style="text-align: center">
      <div class="neumorphic-progress">
        <div class="neumorphic-progress__back"></div>
        <div class="neumorphic-progress__line"></div>
      </div>
      <p>Meter</p>
    </v-col>
    <v-col>
      <NuxtLink to="route/Select">
        <a class="neumorphic-btn mt-12 ml-0">Get dopamine</a>
      </NuxtLink>
    </v-col>
    <Celebration v-if="true"></Celebration>
    <!-- <v-btn nuxt to="route/Select">select</v-btn>
    </v-col> -->
  </v-row>
</template>

<script>
export default {
  setup() {},

  mounted() {
    navigator.geolocation.getCurrentPosition(this.success, this.error)
  },
  methods: {
    success(pos) {
      console.log(pos)
    },
    error(err) {
      console.log(err)
    },
  },
}
</script>
<style scoped>
* {
  text-decoration: none;
}
.up-btn {
  position: relative;
  display: inline-block;
  text-decoration: none;
  padding: 10px 30px;
  text-transform: uppercase;
  font-weight: 500;
  letter-spacing: 2px;
  color: #5a84a2;
  background: #ffffff;
  font-size: 18px;
  border-radius: 40px;
  box-shadow: -2px -2px 8px rgba(255, 255, 255, 1),
    -2px -2px 12px rgba(255, 255, 255, 0.5),
    inset 2px 2px 4px rgba(255, 255, 255, 0.1), 2px 2px 8px rgba(0, 0, 0, 0.15);
}

.down-btn {
  position: relative;
  display: inline-block;
  text-decoration: none;
  padding: 10px 30px;
  text-transform: uppercase;
  font-weight: 500;
  letter-spacing: 2px;
  color: #5a84a2;
  font-size: 18px;
  border-radius: 40px;
  box-shadow: -2px -2px 8px rgba(255, 255, 255, 1),
    -2px -2px 12px rgba(255, 255, 255, 0.5),
    inset 2px 2px 4px rgba(255, 255, 255, 0.1), 2px 2px 8px rgba(0, 0, 0, 0.15);
}
/*hover effect*/
.up-btn:hover {
  box-shadow: inset -2px -2px 8px rgba(255, 255, 255, 1),
    inset -2px -2px 12px rgba(255, 255, 255, 0.5),
    inset 2px 2px 4px rgba(255, 255, 255, 0.1),
    inset 2px 2px 8px rgba(0, 0, 0, 0.15);
}

.up-btn:active {
  background: #e8e8e8;
  box-shadow: inset 3px 3px 8px #bebebe, inset -3px -3px 8px #ffffff;
}

.btn-frame {
  width: 220px;
  background: #e8e8e8;
  box-shadow: inset 3px 3px 8px #bebebe, inset -3px -3px 8px #ffffff;
  border-top-left-radius: 25px;
  border-top-right-radius: 25px;
  border-bottom-right-radius: 25px;
  border-bottom-left-radius: 25px;
}

.btn {
  width: 200px;
  height: 100px;
  border-radius: 50%;
  background: #55b9f3;
  box-shadow: 35px 35px 70px #bebebe, -35px -35px 70px #bebebe;
}
.btn:hover {
  border-radius: 119px;
  background: #e8e8e8;
  box-shadow: inset 26px 26px 47px #bebebe, inset -26px -26px 47px #bebebe;
}
h1 {
  font-family: 'Roboto', sans-serif;
  text-align: center;
  color: #747474;
  text-shadow: 4px 4px 3px #bebebe, -3px -3px 5px #ffffff;
  font-size: 40px;
  letter-spacing: 0.2em;
  text-indent: 0.2em;
  font-weight: 800;
}
p {
  font-family: 'Roboto', sans-serif;
  text-align: center;
  color: #747474;
  text-shadow: 4px 4px 3px #bebebe, -3px -3px 5px #ffffff;
  font-size: 35px;
  letter-spacing: 0.2em;
  text-indent: 0.2em;
  font-weight: 800;
}
.neumorphic-progress,
.neumorphic-slider {
  width: 260px;
  margin: auto;
}
.neumorphic-progress__back,
.neumorphic-slider__back {
  height: 24px;
  background-color: var(--back-color);
  border-radius: 10px;
  border: 4px solid #f3f4f7;
  box-shadow: 7px 7px 15px rgba(55, 84, 170, 0.15),
    -7px -7px 20px rgba(255, 255, 255, 1),
    inset 0px 0px 4px rgba(255, 255, 255, 0),
    inset 7px 7px 15px rgba(55, 84, 170, 0.15),
    inset -7px -7px 20px rgba(255, 255, 255, 1),
    0px 0px 4px rgba(255, 255, 255, 0.2) !important;
}
.neumorphic-slider {
  width: 240px;
}
.neumorphic-slider__back {
  width: calc(100% + 20px);
  margin-left: -10px;
}
.neumorphic-slider {
  position: relative;
}
.neumorphic-progress__line {
  width: var(--value);
  height: 16px;
  background-color: #185bf1;
  margin-top: -20px;
  margin-left: 4px;
  border-radius: 8px;
  opacity: 1;
  animation: sliding 3s ease infinite;
}

@keyframes sliding {
  0% {
    width: 0;
  }
  100% {
    width: calc(100% - 8px);
  }
}

.neumorphic-slider__line {
  height: 16px;
  background-color: #175bec;
  margin-top: -20px;
  margin-left: -6px;
  border-radius: 8px;
  opacity: 1;
  width: calc(var(--value) + 4px);
}

.neumorphic-slider__popover {
  position: absolute;
  background-color: var(--back-color);
  left: var(--value);
  width: 60px;
  height: 40px;
  top: -60px;
  margin-left: -30px;
  box-shadow: 7px 7px 15px rgba(55, 84, 170, 0.15),
    -7px -7px 20px rgba(255, 255, 255, 1), 0px 0px 4px rgba(255, 255, 255, 0.2);
  color: #6d7587;
  padding: 8px 0;
}

.neumorphic-slider__popover::after {
  content: '';
  position: absolute;
  width: 0;
  height: 0;
  bottom: -6px;
  left: 50%;
  margin-left: -6px;
  border-style: solid;
  border-width: 6px 6px 0 6px;
  border-color: var(--back-color) transparent transparent transparent;
}

.neumorphic-slider__thumb {
  box-shadow: 7px 7px 15px rgba(55, 84, 170, 0.15),
    0px 0px 15px rgba(55, 84, 170, 0.3);
  width: 20px;
  height: 20px;
  position: absolute;
  border-radius: 50%;
  background-color: var(--back-color);
  left: var(--value);
  top: 2px;
  margin-left: -10px;
  cursor: pointer;
}

.btn:focus {
  outline: none !important;
}

.neumorphic-btn {
  text-decoration: none;
  align-items: center;
  text-align: center;
  font-size: 45px;
  padding: 0px 40px;
  border-radius: 50%;
  font-weight: bold;
  width: 350px;
  height: 350px;
  color: #2b2b2b;
  display: flex;
  background-color: var(--back-color);
  /*box-shadow: 3px 4px 6px 1px rgba(0, 0, 0, 0.3),
              -2px -2px 4px 2px rgba(255, 255, 255, 1),
              inset -14px -14px 20px -10px rgba(0, 0, 0, 0.1),
              inset 14px 14px 20px -10px rgba(255, 255, 255, 0.5) !important;  */
  box-shadow: 7px 7px 15px rgba(55, 84, 170, 0.15),
    -7px -7px 20px rgba(255, 255, 255, 1),
    inset 0px 0px 4px rgba(255, 255, 255, 0.2),
    inset 7px 7px 15px rgba(55, 84, 170, 0),
    inset -7px -7px 20px rgba(255, 255, 255, 0),
    0px 0px 4px rgba(255, 255, 255, 0) !important;
  transition: box-shadow 0.25s ease !important;
}

.neumorphic-btn:active {
  box-shadow: 7px 7px 15px rgba(55, 84, 170, 0.15),
    -7px -7px 20px rgba(255, 255, 255, 1),
    inset 0px 0px 4px rgba(255, 255, 255, 0),
    inset 7px 7px 15px rgba(55, 84, 170, 0.15),
    inset -7px -7px 20px rgba(255, 255, 255, 1),
    0px 0px 4px rgba(255, 255, 255, 0.2) !important;
}
</style>
