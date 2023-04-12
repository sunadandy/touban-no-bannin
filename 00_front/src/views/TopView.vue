<template>
    <v-main class="top">
        <router-link to="/home" @click="RingAudio()">当番の番人</router-link>
        <br>※クリック音が鳴ります
    </v-main>
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
        }
    }
}
</script>

<style scoped>
h2 {
text-align: center;
}
.app {
width: 900px;
margin: auto;
}
</style>