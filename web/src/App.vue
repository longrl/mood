<template>
  <div class="content">
    <div class="header">
      <divc class="header-title">
        <a href="/">talk is cheap show me the code</a>
      </divc>
      <div class="header-subtitle">keep the love, never forget the beginning</div>
      <div class="header-items">
        <div class="item"><a href="/">About</a></div>
        <div class="item"><a href="https://github.com/longrl/mood">Github</a></div>
        <div class="item" @click="link()">*</div>
      </div>
      <div class="header-line"></div>
    </div>
    <router-view></router-view>
    <div class="bottom">© 湘ICP备123456789号</div>
  </div>

</template>

<script>
  export default {
    methods: {
      link() {
        this.$router.push({ path: "/blog/update/0"})
      }
    }
  }

</script>

<style scoped>
@import "./assets/font.css";
  body {
    margin: 0px;
    font-family: noto;
  }
  .content {
    max-width: 800px;
    margin: 0px auto;
    padding-top: 32px;
  }

  .header-title {
    font-size: 2rem;
    font-weight: 700;

    padding-top: 32px;
    font-family: ragflij,serif;
  }

  a {
    text-decoration: none;
    color: #000000;
  }
  a:hover {
    background-color: #ffffff;
  }
  .header-subtitle {
    margin-top: 15px;
    color: #666;
  }
  .header-items {
    display: flex;
    justify-content: flex-end;
    text-align: end;
    margin: 10px 8px 10px 8px;
  }
  .item {
    margin: 0px 5px;
    font-family: black;
  }

  .header-line {
    width: 100%;
    border-width: 2px;
    border-color: #482936;
    border-style: solid none none none;
  }

  .bottom {
    margin-top: 70px;
    font-family: black;
    font-size: 0.6rem;
    font-weight: 600;
  }
</style>
