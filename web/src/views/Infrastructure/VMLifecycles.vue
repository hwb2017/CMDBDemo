<template>
  <div :style="{ padding: '24px', background: '#fff', minHeight: '360px', margin: '24px 8px'}">
    <a-card
      style="width:100%"
      :title="`${item.Applicant}的主机申请(${item.CreationTime})`"
      :tab-list="tabList"
      :active-tab-key="key"
      @tabChange="key => onTabChange(key, 'key')"
      v-for="item in vmLifecycles"
      :key="item._id"
    >
      <a slot="extra" href="#">More</a>
      <div v-if="key == 'tab1'">
        <a-descriptions layout="vertical">
          <a-descriptions-item label="申请人">
            {{ item.Applicant }}
          </a-descriptions-item>
          <a-descriptions-item label="维护者">
            {{ item.Maintainer }}
          </a-descriptions-item>
          <a-descriptions-item label="申请到期时间">
            {{ item.VMLifecycleRules[0].actiontime }}
          </a-descriptions-item>         
        </a-descriptions>
      </div>
      <div v-else>
        关联主机信息...
      </div>
    </a-card>
  </div>
</template>

<script>
import { mapActions } from 'vuex';  
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
      key: 'tab1',
    };
  },
  computed: {
    vmLifecycles() {
      return this.$store.state.infra.vmLifecycles
    }
  },
  methods: {
    onTabChange(key, type) {
      this[type] = key;
    },
    ...mapActions({
      'getVMLifecycles': 'infra/getVMLifecycles'
    })
  },
  mounted: function() {
    this.getVMLifecycles();
  }
};
</script>

<style lang="less">
</style>