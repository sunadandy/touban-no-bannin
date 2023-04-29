<template>
  <div>
    <ul class="dd-box"
      v-for="(data, index) in dataList" :key="index"
      v-bind:class="index === dragIndex ? 'dragging' : ''"
    >
      <li
      draggable="true"
      @dragstart="dragStart($event, index)"
      @drop.stop="dragStop($event, index)"
      @dragenter="dragEnter($event, index)"
      @dragover.prevent
      >
      {{ data.name }}
      </li>
    </ul>
  </div>
</template>
  
<script>
export default {
  name: 'DragDropSwap',
  data(){
    return{
      dragIndex: null,
      // dataList: this.memberList,
    }
  },
  props:{
    dataList:{
      type: Array,
      required: true,
      default: () => [],
    }
  },
  methods:{
    dragStart(event, index){
      this.dragIndex = index
    },
    dragEnter(event, index){
      if (index === this.dragIndex) return;
      const deleteList = this.dataList.splice(this.dragIndex, 1);  //指定したインデックスから長さ1の要素を削除
      this.dataList.splice(index, 0, deleteList[0])
      this.dragIndex = index
    },
    dragStop(event, index){
      this.dragIndex = null
    },
  }
};
</script>

<style scoped>
.dd-box {
    list-style: none;
    padding: 2px;
}
.dd-box li {
    user-select: none;
    box-shadow: 0 0 2px 1px rgba(0, 0, 0, .2);
    cursor: grab;
    transition: box-shadow .3s;
    padding: 3px 2em;
    font-size: 13px;
    display: table-cell;
    vertical-align: middle;
}
.dragging {
  background-color: #b4b4b4;
}
</style>