<template>
  <div class="blog">
    <v-md-editor
        v-model="data.content"
        height="500px"
        :disabled-menus="[]"
        @upload-image="handleUploadImage"
    ></v-md-editor>
  </div>
  <div class="input">
    <div></div>
    <div class="music_url">
      <div class="music_url_input">
        <el-input v-model="data.music_url" disabled placeholder="music url"/>
      </div>
      <el-upload
          class="upload-demo"
          action="http://localhost:8080/blog/upload"
          show-file-list= false
          :on-success="upload"
      >
        <el-button type="info">Click to upload</el-button>
      </el-upload>
    </div>
    <div class="submit">
      <div v-if="flag == 0" class="create" @click="create"><a>Create</a></div>
      <div v-if="flag == 1" class="update" @click="update"><a>Update</a></div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  const http = axios.create({
    baseURL: "http://localhost:8080/blog",
    timeout: 120000
  });
  export default {
    created() {
      const id = this.$route.params.id

      if (id === "0") {
        this.flag = 0
      } else {
        console.log(id)
        http.get("detail/" + id).then(({data}) => {
          this.data = data.data
        })
        this.flag = 1
      }
    },
    data() {
      return {
        data: {
          "id": "",
          "title": "测试",
          "top": "",
          "image": "",
          "content": "",
          "md_path": "",
          "music_url": "",
          "category": ""
        },

        flag: 0
      }
    },
    methods: {
      upload(response) {
        console.log(response.url)
        this.data.music_url = response.url
      },
      create() {
        console.log(this.data)
        http.post("add", this.data, {headers : {
          'Content-Type' : 'application/json; charset=UTF-8'
        }}).then(({data}) => {
          console.log(data.data)
        })
      },
      update() {

      },
      handleUploadImage(event, insertImage, files) {
        // 拿到 files 之后上传到文件服务器，然后向编辑框中插入对应的内容
        let formdata = new FormData()
        formdata.append('file', files[0])
        console.log(files)
        http.post("upload", formdata, {headers : {'Content-Type': 'multipart/form-data'}}).then(({data}) => {
          insertImage({
            url: data.url,
            width: "200px",
            height: "200px"
          });
        })
      },
    }
  }
</script>

<style scoped>
  .input {
    margin-top: 30px;
  }
  .music_url {
    display: flex;
    justify-content: space-between;
  }
  .music_url_input {
    width: 380px;
  }
  .submit {
    margin-top: 30px;
    display: flex;
  }
  .create {
    border: 2px solid #332e2e;
    border-radius: 5px;
    font-size: 1rem;
    padding: 2px;
    font-family: black;
  }
  .update {
    border: 2px solid #332e2e;
    border-radius: 5px;
    font-size: 1rem;
    padding: 2px;
    font-family: black;
  }
  .update:hover {
    background-color: #d9d6d6;
  }
  .create:hover {
    background-color: #d9d6d6;
  }
</style>