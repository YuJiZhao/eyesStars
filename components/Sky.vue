<template>
  <canvas id="canvas"></canvas>
</template>

<script setup lang="ts">
import * as zrender from "zrender";
import { drawProvision, drawPermanent } from "@/components/paint/script";

const monitor = useMonitor();
const contextStore = useContext();

// 初始化画布
function initCanvas() {
  let canvas = <HTMLCanvasElement>document.getElementById("canvas");
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
  drawCanvas(canvas);
}

// 绘制画布
function drawCanvas(canvas: HTMLCanvasElement) {
  let zr = zrender.init(canvas);
  drawProvision(zr, monitor.value, contextStore.value);
  drawPermanent(zr, monitor.value, contextStore.value);
}

onMounted(() => {
  initCanvas();
});
</script>

<style scoped>
canvas {
  position: fixed;
  left: 0;
  top: 0;
}
</style>