<template>
  <div class="top">
    <div class="inner">
      <img src="./../assets/logo.jpg" @click="RingAudio()">
    </div>
    <div class="caution">画像をクリックすると次に進みます。(なお、音が鳴るので注意！)</div>
  </div>
  <button type="button" @click="PostRequest">追加</button><br>
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
      this.audio.play() // 鳴らす
      this.$router.push("/home")
    },
    PostRequest(){
      // var data = {"title": "sample4", "owner": "yamada", "start": "2023-04-01", "mail": false}
      var data = [
        {"toubanID": 0, "name": "yamada", "employeeNo": "11111", "order": 1, "next": "2023-05-01", "last": "2023-04-01"},
        {"toubanID": 0, "name": "sato", "employeeNo": "22222", "order": 2, "next": "2023-05-01", "last": "2023-04-01"},
        {"toubanID": 1, "name": "yamada", "employeeNo": "11111", "order": 1, "next": "2023-05-01", "last": "2023-04-01"},
      ]
      // this.axios.post("/touban", data)
      this.axios.post("/order", data)
      .then(response => {
        console.log(response)
      }).catch(error => console.log(error))
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