<template>
  <div :class="'tip tip_' + tipClass">
    <transition-group name="tipAnimate">
      <div class="word" v-if="show">{{ msg }}</div>
    </transition-group>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { LayoutModeEnum } from "@/constant/enum";
import config from "@/config";

let props = defineProps({
  show: { type: Boolean },
  msg: { type: String },
});

const monitorStore = useMonitor();
const processStore = useProcess();
let tipClass = ref(monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip");

watch(
  () => props.show,
  (value) => {
    if (value) {
      setTimeout(() => {
        processStore.value.isShowTip = false;
      }, config.tipDuration);
    }
  }
);
</script>

<style scoped lang="scss">
@import "@/assets/scss/global.scss";

.tipAnimate-enter-active {
  animation: fadeIn 0.8s;
}
.tipAnimate-leave-active {
  animation: fadeOut 0.8s;
}

.tip {
  position: fixed;
  z-index: $zIndexTip;
  .word {
    background: #fff;
    border-radius: 5px;
    box-shadow: 0 0 5px rgba($color: #fff, $alpha: 0.7);
    -webkit-box-shadow: 0 0 5px rgba($color: #fff, $alpha: 0.7);
    -moz-box-shadow: 0 0 5px rgba($color: #fff, $alpha: 0.7);
    padding: 5px 10px;
  }
}

.tip_normal {
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  .word {
    max-width: 300px;
  }
}

.tip_flip {
  top: 50%;
  right: -50px;
  transform: translateY(-50%) rotateZ(90deg);
  .word {
    min-width: 150px;
    max-width: 200px;
    text-align: center;
  }
}
</style>