<template>
  <div>
    <v-overlay :model-value="overlay" class="align-center justify-center">
    <v-progress-circular
        color="primary"
        indeterminate
        size="64"
    ></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
export default {
  name: 'Overlay',
  data(){
    return {
      overlay: false,
    }
  },
  props:{
    waitTime: {
      type: Number,
      required: true,
    }
  },
  watch: {
    overlay (val) {
      val && setTimeout(() => {
        this.overlay = false
      }, this.waitTime)
    },
  },
  methods:{
    SetOverlay(){
      this.overlay = true
    }
  }
}
</script>