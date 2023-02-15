<template>
  <canvas id="canvas"></canvas>
</template>

<script setup lang="ts">
import * as zrender from "zrender";
import { drawDynamic, drawStatic } from "@/components/stars/factory";

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
  zr.add(drawDynamic());
  zr.add(drawStatic());
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