<template>
  <canvas id="canvas"></canvas>
</template>

<script setup lang="ts">
import { init } from "zrender";
import { registerPainter } from 'zrender';
import CanvasPainter from 'zrender/lib/canvas/Painter';
import SVGPainter from 'zrender/lib/svg/Painter';
import { drawProvision, drawPermanent } from "@/bussiness/paint/script";

registerPainter('canvas', CanvasPainter);
registerPainter('svg', SVGPainter);

const monitorStore = useMonitor();
const contextStore = useContext();
const processStore = useProcess();

// 初始化并绘制画布
function drawCanvas() {
  let canvas = <HTMLCanvasElement>document.getElementById("canvas");
  let zr = init(canvas, {
    width: window.innerWidth,
    height: window.innerHeight
  });
  drawProvision(zr, monitorStore.value, processStore.value);
  drawPermanent(zr, monitorStore.value, contextStore.value, processStore.value);
}

onMounted(() => {
  drawCanvas();
});
</script>

<style scoped>
canvas {
  position: fixed;
  left: 0;
  top: 0;
}
</style>