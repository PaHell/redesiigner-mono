<template>
      <div class="app-page">
            <div>
                  <slot/>
            </div>
      </div>
</template>

<style global lang="postcss">
	.app-page {
            @apply flex justify-center;

            & > div {
                  @apply w-full max-w-screen-xl;
            } 
	}
</style>
