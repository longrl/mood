<template>
  <div class="music" id="top">
    <div class="img" @click="play()"></div>
    <div class="line inner" id="step"></div>
    <audio :src="data.music_url" id="music"></audio>
  </div>
  <div class="blog">
    <v-md-editor v-model="data.content" mode="preview"></v-md-editor>
  </div>
  <div class="editor">
    <a @click="link(data.id)">editor</a>
  </div>
  <div class="top">
    <a href="#top">top</a>
  </div>
</template>

<script>
  import axios from 'axios'
  const http = axios.create({
    baseURL: "http://localhost:8080/blog",
    timeout: 120000
  });
  export default {
    data() {
      return {
        data: {},
        isPlay: false,
      }
    },
    created() {
      const id = this.$route.params.id
      http.get("detail/" + id).then(({data}) => {
        this.data = data.data
      })
    },
    methods: {
      play() {
        if (!this.isPlay) {
          music.play()
          this.isPlay = true
        } else {
          music.pause()
          this.isPlay = false
        }
        music.addEventListener("timeupdate", function() {
          var step = (music.currentTime / music.duration) * 100
          document.getElementById('step').style.width = step + '%'
          console.log(step + '%')
        });
      },
      link(id) {
        this.$router.push({ path: "/blog/update/" + id })
      }
    }
  }
</script>

<style scoped>
  .blog {
    font-family: noto,serif;
  }

  .editor {
    float: right;
    margin: 20px;
    border: 2px solid #332e2e;
    border-radius: 5px;
    font-size: 0.8rem;
    padding: 2px;
    font-family: black;
  }

  .top {
    float: right;
    margin: 20px;
    border: 2px solid #332e2e;
    border-radius: 5px;
    font-size: 0.8rem;
    padding: 2px;
    font-family: black;
  }

  .editor:hover {
    background-color: #d9d6d6;
  }
  .top:hover {
    background-color: #d9d6d6;
  }

  a {
    text-decoration: none;
    color: #000000;
  }
  a:hover {
    background-color: #ffffff;
  }

  .music {
    height: 30px;
    background-color: rgba(211, 211, 211, 0.15);
    /*display: flex;*/
  }

  .img {
    margin-left: 3px;
    width: 28px;
    height: 28px;
    background:url("../assets/play1.png") no-repeat;
    background-size:100% 100%;
    float: left;
  }
  .img:hover {
    background:url("../assets/play2.png") no-repeat;
    background-size:100% 100%;
  }

  .inner {
    height: 100%;
    border-radius: 4px;
    transition: all 0.5s cubic-bezier(0, 0.64, 0.36, 1);
  }

  .line {
    width: 0%;
    background-color: #e0e0e0;
  }

</style>