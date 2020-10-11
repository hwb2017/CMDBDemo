<template>
  <div :style="{ padding: '24px', background: '#fff', minHeight: '360px', margin: '24px 8px'}">
    <a-input :placeholder="searchItemPlaceholder" style="width: 320px " v-model="searchItemValue">
      <a-select slot="addonBefore" default-value="ipAddr" v-model="searchItemKey" style="width: 120px" @change="changePlaceholder">
        <a-select-option value="ipAddr">
          IP地址
        </a-select-option>
        <a-select-option value="hostName">
          主机名称
        </a-select-option>
      </a-select>
    </a-input>
    <a-icon type="search" :style="{ fontSize: '28px', color: '#08c', margin: '2px 8px'}" @click="handleSearchItemChange"/>
    <div :style="{margin: '8px 2px'}">
      <span :style="{ margin: '8px 2px' }">分组标签:</span>
      <template v-for="tag in tags">
        <a-checkable-tag
          :key="tag"
          :checked="selectedTags.indexOf(tag) > -1"
          @change="checked => handleTagChange(tag, checked)"
        >
          {{ tag }}
        </a-checkable-tag>
      </template>
    </div>
    <a-table 
      :columns="columns" 
      :data-source="virtualmachines" 
      :row-key="record => record.InstanceID"
      :pagination=false
    >
      <div slot="IpAddress" slot-scope="record">
        <p v-if="record.public_ip_address">{{ record.public_ip_address }} (公网)</p>
        <p>{{ record.private_ip_address }} (内网)</p>
      </div>
    </a-table>
    <a-pagination
      class="pagination"
      v-model="pageNum"
      :defaultPageSize=10
      :total="vmTotal"
      :showTotal="total => `总共 ${total} 项`"
      :pageSize.sync="pageSize"
      :pageSizeOptions="pageSizeOptions"
      showSizeChanger
      @change="handlePaginationChange"
      @showSizeChange="handlePageSizeChange"
    />
  </div>
</template>

<script>
import { mapActions } from 'vuex';
const columns = [
  {
    key: 'Provider',
    dataIndex: 'vm_provider',    
    title: '提供商',
  },
  {
    key: 'InstanceName',
    dataIndex: 'instance_name',
    title: '主机名称',    
  },
  {
    key: 'OSName',
    dataIndex: 'os_name',
    title: '操作系统', 
  },
  {
    key: 'IpAddress',    
    title: 'IP地址',
    scopedSlots: { customRender: 'IpAddress' },
  },
  {
    key: 'InstanceType',
    dataIndex: 'instance_type',
    title: '实例类型'
  }
];
const providerMapping = {
  "阿里云": "alicloud",
  "AWS": "aws"
}
const placeholderMapping = {
  "ipAddr": "IP地址",
  "hostName": "主机名称"
}
export default {
  data() {
    return {
      columns,
      searchItemPlaceholder: "请输入IP地址",
      searchItemValue: "",
      searchItemKey: "ipAddr",
      tags: ['阿里云', 'AWS'],
      selectedTags: [],
      pageSizeOptions: ['10', '20', '50'],
      pageSize: 10,
      pageNum: 1
    };
  },
  computed: {
    virtualmachines() {
      return this.$store.state.infra.virtualmachines
    },
    vmTotal() {
      return this.$store.state.infra.vmTotal
    },
    queryOptions() {
      return {
        pageNum: this.pageNum,
        pageSize: this.pageSize,
        provider: this.selectedTags.map((item) => {
          return providerMapping[item];
        }),
        searchItemKey: this.searchItemKey,
        searchItemValue: this.searchItemValue
      }
    }
  },
  methods: {
    ...mapActions({
      'getVirtualMachines': 'infra/getVirtualMachines'
    }),
    changePlaceholder(value) {
      this.searchItemPlaceholder = `请输入${placeholderMapping[value]}`;
    },
    handleTagChange(tag, checked) {
      const { selectedTags } = this;
      const nextSelectedTags = checked
        ? [...selectedTags, tag]
        : selectedTags.filter(t => t !== tag);
      this.selectedTags = nextSelectedTags;
      this.getVirtualMachines(this.queryOptions);
    },
    handlePaginationChange() {
        this.getVirtualMachines(this.queryOptions);
    },
    handlePageSizeChange() {
        this.getVirtualMachines(this.queryOptions);
    },
    handleSearchItemChange() {
        this.getVirtualMachines(this.queryOptions);
    }
  },
  mounted: function() {
    var pagination = {
      pageNum: 1,
      pageSize: 10
    }
    this.getVirtualMachines(pagination);
  }
};  
</script>

<style lang="less">
.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>