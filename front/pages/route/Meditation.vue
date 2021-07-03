<template>
  <body>
    <center class="all">
      <p color="black" class="texttmp">何分間瞑想しますか？</p>
      <v-row>
        <v-col>
          <v-btn
            large
            elevation="4"
            height="5.5rem"
            width="12rem"
            style="text-transform: none"
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
          >
            <div class="font">8min</div>
          </v-btn>
        </v-col>
      </v-row>
    </center>
  </body>
</template>

<script>
export default {}
</script>

<style>
.texttmp {
  color: black;
}
.all {
  margin-top: 10px;
  font-size: 7.7vh;
}
.font {
  font-size: 32px;
  font-weight: bold;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
</style>