<script setup>
import Dialog from "/node_modules/primevue/dialog";
</script>

<template>
  <Dialog v-model:visible="display">
    Hello World!
  </Dialog>
</template>

<script>
export default {
  props: ['display'],
}

//export default {
//  data() {
//    return {
//      display: true
//    }
//  }
//}
</script>
