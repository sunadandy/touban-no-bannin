<template>
  <div class="top">
    <div class="inner">
      <img src="./../assets/logo.jpg" @click="RingAudio()">
    </div>
    <div class="caution">画像をクリックすると次に進みます。(なお、音が鳴るので注意！)</div>
  </div>
</template>

<script>
export default {
  name: 'TopView',
  data: () => ({
  audio: new Audio(require('@/assets/バーン.mp3'))
  }),
  methods: {
    RingAudio: function() {
      this.audio.currentTime = 0 // 連続で鳴らせるように
      // this.audio.play() // 鳴らす
      this.$router.push("/home")
    },
  }
}
</script>

<style scoped>
.top{
  height: 400px;
  position: relative;
  text-align: center;
  /* background: #ffb21d; */
}
.inner{
  left: 50%;
  position: absolute;
  top: 50%;
  transform: translate(-50%,-50%);
}
.caution{
  left: 50%;
  position: absolute;
  top: 70%;
  transform: translate(-50%,-50%);
}
</style>