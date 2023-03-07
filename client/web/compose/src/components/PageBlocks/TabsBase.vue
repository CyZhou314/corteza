<template>
  <wrap
    v-bind="$props"
    :scrollable-body="false"
    v-on="$listeners"
  >
    <div
      v-if="!options.tabs.length"
      class="d-flex h-100 align-items-center justify-content-center"
    >
      <p class="mb-0">
        {{ $t('tabs.noTabs') }}
      </p>
    </div>

    <b-tabs
      v-else
      card
      nav-class="bg-white"
      :nav-wrapper-class="navWrapperClass"
      :content-class="contentClass"
      v-bind="{
        align: block.options.style.alignment,
        fill: block.options.style.fillJustify === 'fill',
        justified: block.options.style.fillJustify === 'justify',
        pills: block.options.style.appearance === 'pills',
        tabs: block.options.style.appearance === 'tabs',
        small: block.options.style.appearance === 'small',
        vertical: block.options.style.orientation === 'vertical',
        end: block.options.style.position === 'end'
      }"
      lazy
      class="h-100"
      :class="{ 'd-flex flex-column': block.options.style.orientation !== 'vertical' }"
    >
      <b-tab
        v-for="(tab, index) in tabbedBlocks"
        :key="index"
        :title="getTabTitle(tab, index)"
        class="h-100"
        :title-item-class="titleItemClass"
        :title-link-class="titleItemClass"
        no-body
      >
        <page-block-tab
          v-if="tab.block"
          v-bind="{ ...$attrs, ...$props, page, block: tab.block, blockIndex: index }"
          :record="record"
          :module="module"
        />

        <div
          v-else
          class="d-flex h-100 align-items-center justify-content-center"
        >
          <p class="mb-0">
            {{ $t('tabs.noBlock') }}
          </p>
        </div>
      </b-tab>
    </b-tabs>
  </wrap>
</template>

<script>
import base from './base'
import { compose } from '@cortezaproject/corteza-js'
import { fetchID } from 'corteza-webapp-compose/src/lib/tabs'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'TabBase',

  components: {
    PageBlockTab: () => import('corteza-webapp-compose/src/components/PageBlocks'),
  },

  extends: base,

  computed: {
    tabbedBlocks () {
      return this.block.options.tabs.map(({ blockID, title }) => {
        const block = this.page.blocks.find(b => fetchID(b) === blockID)
        return {
          block: block ? compose.PageBlockMaker(block) : undefined,
          title,
        }
      })
    },

    contentClass () {
      return `card overflow-hidden ${this.block.options.style.orientation === 'vertical' ? 'd-block' : 'flex-fill'}`
    },

    navWrapperClass () {
      const { orientation, position } = this.block.options.style
      let border = 'border-bottom'

      if (orientation === 'vertical') {
        border = position === 'end' ? 'border-left' : 'border-right'
      } else if (position === 'end') {
        border = 'border-top'
      }

      return `bg-white ${border}`
    },

    titleItemClass () {
      const { fillJustify, alignment } = this.block.options.style
      return `text-truncate text-${alignment} ${fillJustify !== 'none' ? 'flex-fill' : ''}`
    },
  },

  methods: {
    getTabTitle ({ title, block = {} }, tabIndex) {
      const { title: blockTitle, kind } = block
      return title || blockTitle || kind || `${this.$t('tabs.tab')} ${tabIndex + 1}`
    },
  },
}
</script>
