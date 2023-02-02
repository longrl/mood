<template>
  <div class="main">
    <div v-for="item in blog">

      <div class="blog">
        <div class="item-date">{{item.date}}</div>
        <div class="item-title" @click="link(item.id)">{{item.title}}</div>
        <div class="item-category">#{{category[item.category]}}</div>
      </div>
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
    data() {
      return {
        blog: [],
        category:["", "技术", "随想"]
      }
    },
    created() {
      http.get("list").then(({data}) => {
        this.blog = data.data
      })
    },
    methods: {
      link(id) {
        this.$router.push({ path: "/blog/" + id })
      }
    }
  }
</script>

<style scoped>
  .main {
    max-width: 800px;
    font-size: 1rem;
    font-weight: 500;
    font-family: noto;
    margin-top: 40px;
  }

  .blog {
    display: flex;
    justify-content: space-between;
    margin-bottom: 10px;
  }
  .item-category {
    font-family: black;
    font-size: 0.6rem;
    font-weight: 600;
  }

  .item-date {
    font-family: black;
    font-weight: 600;
  }

  .item-title:hover {
    font-weight: 600;
  }
</style>