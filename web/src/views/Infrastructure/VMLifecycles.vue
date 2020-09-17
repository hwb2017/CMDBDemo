<template>
  <div :style="{ padding: '24px', background: '#fff', minHeight: '360px', margin: '24px 8px'}">
    <a-card
      style="width:100%"
      title="Carol的主机申请-20200917"
      :tab-list="tabList"
      :active-tab-key="key"
      @tabChange="key => onTabChange(key, 'key')"
    >
      <a slot="extra" href="#">More</a>
      {{ contentList[key] }}
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tabList: [
        {
          key: 'tab1',
          tab: '基本信息',
        },
        {
          key: 'tab2',
          tab: '关联主机',
        },
      ],
      contentList: {
        tab1: '申请人：Carol         维护者：Kevin\n申请到期时间：2020年10月8日',
        tab2: '关联主机信息...',
      },
      key: 'tab1',
    };
  },
  methods: {
    onTabChange(key, type) {
      this[type] = key;
    },
  },
};
</script>

<style lang="less">
</style>