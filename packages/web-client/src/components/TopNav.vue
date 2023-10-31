<template>
  <div class="topNav">
    <div class="topLeft">
      <!-- <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-31liebiao"></use>
      </svg> -->
    </div>

    <div class="topContent">
      <span>直播</span>
      <span>热点</span>
      <span>同城</span>
      <span class="active">推荐</span>
    </div>

    <div class="topRight">
      <svg class="icon" aria-hidden="true" @click="$router.push('/search')">
        <use xlink:href="#icon-search"></use>
      </svg>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="less" scoped>
.topNav {
  background-color: #161823;
  color: #ffffff;
  width: 100%;
  height: 1rem;
  padding: 0.2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .topRight {
    .icon {
      fill: #ffffff;
      width: 0.5rem;
      height: 0.5rem;
    }
  }
  // less中可以样式包样式(div中也是包裹的)
  .topContent {
    width: 65%;
    height: 100%;
    display: flex;
    justify-content: space-around;
    align-items: center;
    font-size: 0.36rem;
    .active {
      border-bottom: 2px solid #fff;
      font-weight: 700;
    }
  }
}
</style>
