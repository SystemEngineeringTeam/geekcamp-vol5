<template>
  <v-row>
    <v-col cols=12>
      <div class="black--text text-h4">音楽を聞く</div>
      <v-text-field
        v-model="message"
        style="color:black"
        label="Message"
        type="text"
      >
        <template v-slot:append-outer>
          <v-btn @click="searchMovies()">
            検索
          </v-btn>
        </template>
      </v-text-field>
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
            key: 'AIzaSyDtuov7ul5Wy5rMp37DbbG3Z4uzumM8Qd4'   // (5)
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