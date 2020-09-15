<template>
  <div :style="{ padding: '24px', background: '#fff', minHeight: '360px', margin: '24px 8px'}">
    <a-table :columns="columns" :data-source="virtualmachines" :row-key="record => record.InstanceID">
      <span slot="Provider">阿里云</span>
      <div slot="IpAddress" slot-scope="record">
        <p v-if="record.PublicIpAddress">{{ record.PublicIpAddress }} (公网)</p>
        <p>{{ record.PrivateIpAddress }} (内网)</p>
      </div>
    </a-table>
  </div>
</template>

<script>
import { mapActions } from 'vuex';
const columns = [
  {
    key: 'Provider',    
    title: '提供商',
    scopedSlots: { customRender: 'Provider' },
  },
  {
    key: 'InstanceName',
    dataIndex: 'InstanceName',
    title: '主机名称',    
  },
  {
    key: 'OSName',
    dataIndex: 'OSName',
    title: '操作系统', 
  },
  {
    key: 'IpAddress',    
    title: 'IP地址',
    scopedSlots: { customRender: 'IpAddress' },
  },
  {
    key: 'InstanceType',
    dataIndex: 'InstanceType',
    title: '实例类型'
  }
];
export default {
  data() {
    return {
      columns,
    };
  },
  computed: {
    virtualmachines() {
      return this.$store.state.infra.virtualmachines
    }
  },
  methods: {
    ...mapActions({
      'getVirtualMachines': 'infra/getVirtualMachines'
    })
  },
  mounted: function() {
    this.getVirtualMachines();
  }
};  
</script>

<style lang="less">
</style>