<template>
  <b-card
    class="shadow-sm"
    data-test-id="card-user-info"
    header-bg-variant="white"
    footer-bg-variant="white"
  >
    <template #header>
      <h3 class="m-0">
        {{ $t('title') }}
      </h3>
    </template>

    <b-form
      enctype="multipart/form-data"
      @submit.prevent="$emit('submit', user)"
    >
      <b-form-group
        :label="$t('preview.label')"
        label-cols="2"
      >
        <div class="d-flex justify-content-between">
          <h6>
            {{ $t("preview.title") }}
          </h6>
          <b-button
            v-if="uploadedAvatar('avatar')"
            variant="link"
            class="d-flex align-items-top text-dark p-1"
            @click="$emit('resetAttachment', 'avatar')"
          >
            <font-awesome-icon :icon="['far', 'trash-alt']" />
          </b-button>
        </div>

        <c-uploader-with-preview
          :value="uploadedAvatar('avatar')"
          :endpoint="`/users/${user.userID}/avatar`"
          :labels="$t('uploader', { returnObjects: true })"
          class="w-50"
          @upload="$emit('onUpload')"
          @clear="$emit('resetAttachment', 'avatar')"
        />
      </b-form-group>

      <div class="form-row mt-3">
        <b-form-group
          :label="$t('initial.color')"
          label-cols="4"
          class="col"
        >
          <b-form-input
            v-model="user.meta.avatarColor"
            type="color"
            data-test-id="input-handle"
          />
        </b-form-group>

        <b-form-group
          :label="$t('initial.backgroundColor')"
          label-cols="4"
          class="col"
        >
          <b-form-input
            v-model="user.meta.avatarBgColor"
            type="color"
            data-test-id="input-handle"
          />
        </b-form-group>
      </div>
    </b-form>

    <template #footer>
      <c-submit-button
        class="float-right"
        :processing="processing"
        :success="success"
        :disabled="saveDisabled"
        @submit="$emit('submit', user)"
      />
    </template>
  </b-card>
</template>

<script>
import CSubmitButton from 'corteza-webapp-admin/src/components/CSubmitButton'
import CUploaderWithPreview from 'corteza-webapp-admin/src/components/CUploaderWithPreview'

export default {
  name: 'CUserEditorAvatar',

  i18nOptions: {
    namespaces: 'system.users',
    keyPrefix: 'editor.avatar',
  },

  components: {
    CSubmitButton,
    CUploaderWithPreview,
  },

  props: {
    user: {
      type: Object,
      required: true,
    },
  },

  methods: {
    uploadedAvatar (name) {
      const attachmentID = this.user.meta.avatarID

      if (attachmentID !== '0') {
        return (
          this.$SystemAPI.baseURL +
            this.$SystemAPI.attachmentOriginalEndpoint({
              attachmentID,
              kind: 'avatar',
              name,
            })
        )
      }
    },
  },
}
</script>
