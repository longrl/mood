<template>
  <div class="authority">
    <el-input
        v-model="data.key"
        type="password"
        placeholder="Please input password"
        show-password
    />
    <el-input
        v-model="data.answer"
        type="password"
        placeholder="Please input captcha"
        show-password
    />
    <img :src="img">
    <el-button @click="login">授权</el-button>
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
    http.get("captcha").then(({data}) => {
      this.data.id = data.id
      this.img = data.b64
    })
  },
  data() {
    return {
      data: {
        "id": "",
        "key": "",
        "answer": ""
      },
      img: ""
    }
  },
  methods: {
    login() {
      http.post("authority", this.data, {headers : {
          'Content-Type' : 'application/json; charset=UTF-8'
        }}).then(({data}) => {
        if (data.code !== 1000) {
          this.$notify.error({
            title: "失败",
            message: data.message
          });
        }
      })
    }
  }
}
</script>

<style scoped>
  .authority {
    margin: 20px;
  }
</style>