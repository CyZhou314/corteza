<template>
  <div class="py-3">
    <portal to="topbar-title">
      {{ $t('edit.edit') }}
    </portal>

    <portal to="topbar-tools">
      <b-button-group
        v-if="page && page.canUpdatePage"
        size="sm"
        class="mr-1"
      >
        <b-button
          variant="primary"
          class="d-flex align-items-center"
          :to="{ name: 'admin.pages.builder' }"
        >
          {{ $t('label.pageBuilder') }}
          <font-awesome-icon
            :icon="['fas', 'cogs']"
            class="ml-2"
          />
        </b-button>

        <page-translator
          v-if="page"
          :page="page"
          style="margin-left:2px;"
        />

        <b-button
          variant="primary"
          :title="$t('tooltip.view')"
          :disabled="!pageViewer"
          :to="pageViewer"
          class="d-flex align-items-center"
          style="margin-left:2px;"
        >
          <font-awesome-icon
            :icon="['far', 'eye']"
          />
        </b-button>
      </b-button-group>
    </portal>

    <b-container fluid="xl">
      <b-row no-gutters>
        <b-col>
          <b-card
            no-body
            class="shadow-sm"
          >
            <b-form
              class="px-4 py-3"
            >
              <b-row>
                <b-col
                  cols="12"
                  md="6"
                >
                  <b-form-group
                    :label="`${$t('newPlaceholder')} *`"
                    label-class="text-primary"
                  >
                    <input
                      id="id"
                      v-model="page.pageID"
                      required
                      type="hidden"
                    >
                    <b-form-input
                      v-model="page.title"
                      data-test-id="input-title"
                      required
                      :state="titleState"
                      class="mb-2"
                    />
                  </b-form-group>
                </b-col>
                <b-col
                  cols="12"
                  md="6"
                >
                  <b-form-group
                    :label="$t('label.handle')"
                    label-class="text-primary"
                  >
                    <b-form-input
                      v-model="page.handle"
                      data-test-id="input-handle"
                      :state="handleState"
                      class="mb-2"
                      :placeholder="$t('block.general.placeholder.handle')"
                    />
                    <b-form-invalid-feedback :state="handleState">
                      {{ $t('block.general.invalid-handle-characters') }}
                    </b-form-invalid-feedback>
                  </b-form-group>
                </b-col>
              </b-row>

              <b-form-group
                :label="$t('label.description')"
                label-class="text-primary"
              >
                <b-form-textarea
                  v-model="page.description"
                  data-test-id="input-description"
                  :placeholder="$t('edit.pageDescription')"
                  rows="4"
                />
              </b-form-group>

              <b-form-group
                v-if="!isRecordPage"
              >
                <b-form-checkbox
                  v-model="page.visible"
                  data-test-id="checkbox-page-visibility"
                  switch
                >
                  {{ $t('edit.visible') }}
                </b-form-checkbox>
              </b-form-group>

              <!-- <b-form-group
                data-test-id="checkbox-show-sub-pages-in-sidebar"
                class="d-flex"
                switch
              >
                <b-form-checkbox
                  v-model="a"
                >
                  {{ $t('showSubPages') }}
                </b-form-checkbox>
              </b-form-group> -->

              <b-button
                variant="light"
                size="lg"
                class="text-dark"
                @click="showModal = true"
              >
                <!-- <font-awesome-icon
                  :icon="['fas', 'plus']"
                  class="mr-2"
                /> -->
                {{ $t('icon.set') }}
              </b-button>
              <b-modal
                v-model="showModal"
                size="lg"
                ok-only
                :title="$t('icon.configure')"
                :ok-title="$t('label.saveAndClose')"
                @ok="setIcon"
                @show="onOpen"
              >
                <b-form-group
                  class="mb-0"
                  :label="$t('icon.upload')"
                >
                  <uploader
                    :endpoint="endpoint"
                    :accepted-files="['image/*']"
                    :form-data="uploaderFormData"
                    @uploaded="uploadAttachment"
                  />

                  <b-form-group
                    :label="$t('url.label')"
                    class="mt-2"
                  >
                    <b-input
                      v-model="iconUrl"
                      :disabled="isIconSet"
                      :style="{'cursor': isIconSet ? 'not-allowed' : 'default'}"
                    />
                  </b-form-group>
                </b-form-group>

                <hr>

                <!-- list of uploaded icons -->
                <list-loader
                  kind="page"
                  :set.sync="attachments"
                  :namespace="namespace"
                  :disabled="iconUrl"
                  :style="{'cursor': isIconSet ? 'not-allowed' : 'default'}"
                  enable-icon-select
                  disable-preview
                  mode="gallery"
                  class="h-100"
                  @toggle-selected-icon="toggleSelectedIcon"
                />
              </b-modal>
            </b-form>
          </b-card>
        </b-col>
      </b-row>
    </b-container>

    <portal to="admin-toolbar">
      <editor-toolbar
        :back-link="{ name: 'admin.pages' }"
        :hide-delete="hideDelete"
        :hide-save="!page.canUpdatePage"
        :disable-save="disableSave"
        @clone="handleClone()"
        @delete="handleDeletePage"
        @save="handleSave()"
        @saveAndClose="handleSave({ closeOnSuccess: true })"
      >
        <template #delete>
          <b-dropdown
            v-if="showDeleteDropdown"
            data-test-id="dropdown-delete"
            size="lg"
            variant="danger"
            :text="$t('general:label.delete')"
          >
            <b-dropdown-item
              data-test-id="dropdown-item-delete-update-parent-of-sub-pages"
              @click="handleDeletePage('rebase')"
            >
              {{ $t('delete.rebase') }}
            </b-dropdown-item>
            <b-dropdown-item
              data-test-id="dropdown-item-delete-sub-pages"
              @click="handleDeletePage('cascade')"
            >
              {{ $t('delete.cascade') }}
            </b-dropdown-item>
          </b-dropdown>
        </template>
      </editor-toolbar>
    </portal>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import EditorToolbar from 'corteza-webapp-compose/src/components/Admin/EditorToolbar'
import PageTranslator from 'corteza-webapp-compose/src/components/Admin/Page/PageTranslator'
import pages from 'corteza-webapp-compose/src/mixins/pages'
import Uploader from 'corteza-webapp-compose/src/components/Public/Page/Attachment/Uploader'
import { compose, NoID } from '@cortezaproject/corteza-js'
import { handle } from '@cortezaproject/corteza-vue'
import ListLoader from 'corteza-webapp-compose/src/components/Public/Page/Attachment/ListLoader'

export default {
  i18nOptions: {
    namespaces: 'page',
  },

  name: 'PageEdit',

  components: {
    EditorToolbar,
    PageTranslator,
    Uploader,
    ListLoader,
  },

  mixins: [
    pages,
  ],

  props: {
    namespace: {
      type: compose.Namespace,
      required: true,
    },

    pageID: {
      type: String,
      required: true,
    },
  },

  data () {
    return {
      modulesList: [],
      page: new compose.Page(),
      showModal: false,
      attachments: [],
      isPageIcon: false,
    }
  },

  computed: {
    ...mapGetters({
      pages: 'page/set',
    }),

    titleState () {
      return this.page.title.length > 0 ? null : false
    },

    handleState () {
      return handle.handleState(this.page.handle)
    },

    pageViewer () {
      if (this.isRecordPage) {
        return undefined
      }
      const { pageID } = this.page
      return { name: 'page', params: { pageID } }
    },

    isRecordPage () {
      return this.page && this.page.moduleID !== NoID
    },

    hasChildren () {
      return this.pages.some(({ selfID }) => selfID === this.page.pageID)
    },

    disableSave () {
      return [this.titleState, this.handleState].includes(false)
    },

    hideDelete () {
      return this.hasChildren || !this.page.canDeletePage || !!this.page.deletedAt
    },

    showDeleteDropdown () {
      return this.hasChildren && this.page.canDeletePage && !this.page.deletedAt
    },

    endpoint () {
      return this.$ComposeAPI.pageUploadIconEndpoint({
        namespaceID: this.namespaceID,
      })
    },

    namespaceID () {
      return this.namespace.namespaceID ? this.namespace.namespaceID : NoID
    },


    uploaderFormData () {
      return {
        upload: 'icon'
      }
    },

    iconUrl: {
      get () {
        return this.page.config.navItem.icon ? this.page.config.navItem.icon.src : ''
      },

      set (iconUrl) {
        this.page.config.navItem.icon.src = iconUrl
      },
    },

    isIconSet () {
      this.attachments.find(a => a.isPageIcon)
    }
  },

  watch: {
    pageID: {
      immediate: true,
      handler (pageID) {
        if (pageID) {
          this.findPageByID({ namespaceID: this.namespaceID, pageID }).then((page) => {
            this.page = page.clone()
          }).catch(this.toastErrorHandler(this.$t('notification:page.loadFailed')))
        }
      },
    },
  },

  methods: {
    ...mapActions({
      findPageByID: 'page/findByID',
      updatePage: 'page/update',
      deletePage: 'page/delete',
      createPage: 'page/create',
      loadPages: 'page/load',
    }),

    handleSave ({ closeOnSuccess = false } = {}) {
      /**
       * Pass a special tag alongside payload that
       * instructs store layer to add content-language header to the API request
       */
      const resourceTranslationLanguage = this.currentLanguage
      this.updatePage({ namespaceID: this.namespaceID, ...this.page, resourceTranslationLanguage }).then((page) => {
        this.page = page.clone()
        this.toastSuccess(this.$t('notification:page.saved'))
        if (closeOnSuccess) {
          this.$router.push({ name: 'admin.pages' })
        }
      }).catch(this.toastErrorHandler(this.$t('notification:page.saveFailed')))
    },

    handleDeletePage (strategy = 'abort') {
      this.deletePage({ ...this.page, strategy }).then(() => {
        this.$router.push({ name: 'admin.pages' })
      }).catch(this.toastErrorHandler(this.$t('notification:page.deleteFailed')))
    },

    uploadAttachment ({ attachmentID }) {
      // should I check if Icon was already added?
        // I can check by url
      this.attachments.unshift({
        attachmentID: attachmentID,
        isPageIcon: this.isPageIcon
      })
      // debugger
      // still not working after temp fix was reverted
      // this.$ComposeAPI.pageUploadIcon({
      //   namespaceID: this.namespaceID,
      //   icon: file
      // }).then(res => {
      //   debugger
      //   // add icon to gallery of icons
      //     this.attachments.unshift(attachmentID)
      // })
    },

    onOpen () {
      this.$ComposeAPI.pageListIcons({
        namespaceID: this.namespace.namespaceID,
        pageID: this.pageID,
      })
      // .then(icons => {
      //   // not tested
      //   debugger
      //   if (icons) {
          // icons.forEach(a => {
            // check curr selected icon
            // if (a.src === this.page.config.navItem.icon.src) {
              // this.attachments.push({ attachmentID: a.attachmentID, isPageIcon: true })
            //} else {
              // this.attachments.push(a)
            // }
          // })
      //   } else {
      //     this.attachments = []
      //   }
      // })
    },

    setIcon () {
      debugger
      // handle link scenario
      // handle attachment scenario
      let attachmentType = 'link'
      let attachmentSource = this.iconUrl

      if (!this.iconUrl) {
        attachmentType = 'attachment'
        // get chosen icon source
        attachmentSource = this.attachment.find(att => att.isPageIcon).attachmentID
        attachmentSource = { attachmentSource }
      }

      this.$ComposeAPI.pageUpdateIcon({
        namespaceID: this.namespaceID,
        pageID: this.pageID,
        type: attachmentType,
        source: attachmentSource
      }).then(i => {
        this.page.config.navItem = {
          icon: {
            type: i.type,
            src: i.src,
          }
        }
        // this.page.config = {
        //   icon: {
        //     type: i.type,
        //     src: i.src,
        //   }
        // }
      })
    },

    toggleSelectedIcon ({ attachmentID = '' }) {
      debugger
      // test this!
      this.attachments = this.attachments.map(a => {
        if (a.isPageIcon && !a.attachmentID === attachmentID) {
          a.isPageIcon = false
        } else {
          a.isPageIcon = true
        }

        return a
      })
    },
  },
}
</script>
