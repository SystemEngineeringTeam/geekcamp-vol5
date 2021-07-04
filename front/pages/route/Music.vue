<template>
  <v-row>
    <v-col cols=12>
      <p class="mt-5">listen to music</p>
      <input v-model="message" class="neumorphic-input" id="login__input" type="text" />
      <button @click="searchMovies()">
        検索
      </button>
      </v-col>
      <v-col v-for="(video,i) in videos" :key="i" cols="6">
          <!-- <p style="color:black">{{ video.id.videoId }}</p> -->
          <iframe width="100%" height="150" :src='"https://www.youtube.com/embed/" + video.id.videoId' title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
          <p style="color:black">{{ genre[i] }}</p>
      </v-col>
  </v-row>
</template>

<script>
const axios = require('axios');
export default {
    data() {
      return {
        genre:['洋楽','J-POP','ロック','K-POP','アニメ', '歌謡曲', 'ジャズ', 'クラシック'],
        message: '',
        results: null,  // (1)
        keyword: '',    // (2)
        params: {   // (3)
            q: '',  // 検索文字列
            part: 'snippet',
            type: 'video',
            maxResults: '6', // 最大検索数 // (4)
            // キーを設定
            key: 'AIzaSyAkBisrX2WfLhU7ugMjSqEDdSqGWQdljD0'   // (5)
        },
        videos:[],
        videoSrc: 'https://www.youtube.com/embed/'
      }
    },
    created(){
      let count = 0
      const intervalId = setInterval(() => {
        this.request(count)
        count += 1
        if(count >= 8){
          clearInterval(intervalId);
        }
      },200)
    },
    methods: {
      searchMovies() {   // (6)
        this.keyword = ''
        this.keyword = this.message + ' 最新曲';
        this.params.q = this.keyword;   // (7)
        const self = this;
        axios
            .get('https://www.googleapis.com/youtube/v3/search', {params: this.params}) // (8)
            .then(function(res) {
                self.videos = res.data.items
                console.log(self.videos)
            })
            .catch(function(err) {
                console.log(err);
            });
      },
      request(i){
        this.keyword = ''
        this.keyword = this.genre[i];
        this.params.q = this.keyword // (7)
        console.log(this.params)
        const self = this;
        axios
          .get('https://www.googleapis.com/youtube/v3/search', {params: this.params}) // (8)
          .then(function(res) {
              self.videos.push(res.data.items[0])
          })
          .catch(function(err) {
              console.log(err);
          });
      }
    }
}
</script>

<style scoped>
.neumorphic-label {
  color: #868b98;
  display: block;
  text-align: start;
  margin: 25px 10px 5px;
}
.neumorphic-input {
  background-color: var(--back-color);
  border-radius: 10px;
  border: 3px solid #f3f4f7;
  box-shadow: 2px 2px 3px rgba(55, 84, 170, .15),
              inset 0px 0px 4px rgba(255, 255, 255, 0),
              inset 7px 7px 15px rgba(55, 84, 170, .15),
              inset -7px -7px 20px rgba(255, 255, 255, 1),
              0px 0px 4px rgba(255, 255, 255, .2) !important;
  padding: 5px 15px;
  outline: none !important;
  color: #535D74;
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
  padding: 0.1em 0.2em;
  background: #efefef;
  border: none;
  color: #2b2b2b;
  border-radius: .5rem;
  font-size: 1.5rem;
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

</style>