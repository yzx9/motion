function remSize() {
  // 获取设备宽度
  var deviceWidth = document.documentElement.clientWidth || window.innerWidth;
  if (deviceWidth >= 750) {
    deviceWidth = 750;
  }
  if (deviceWidth <= 320) {
    deviceWidth = 320;
  }
  // fontSize 属性 设置或返回文本的字体尺寸
  document.documentElement.style.fontSize = deviceWidth / 7.5 + "px";
  // if Width=750   1rem==100px

  // 设置字体大小
  document.querySelector("body").style.fontSize = 0.3 + "rem";
}
remSize();
// onresize 事件会在窗口或框架被调整大小时发生。
window.onresize = function () {
  // 当窗口发生变化时重新适配
  remSize();
};
