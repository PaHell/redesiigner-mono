<script lang="ts">
</script>

<template>
	<div class="app-loading-spinner">
		<div></div>
		<div></div>
		<div></div>
		<div></div>
	</div>
</template>

<style global lang="postcss">
	.app-loading-spinner {
		@apply h-[10px]
		flex justify-center space-x-[4px];
		& > div {
			@apply w-[2px] h-full
			bg-fuchsia-500 rounded-full;
			animation: dotScaleOpacity 1.7s infinite ease-in-out both;
			&:nth-child(2) { animation-delay: 0.2s; }
			&:nth-child(3) { animation-delay: 0.4s; }
			&:nth-child(4) { animation-delay: 0.6s; }
		}
	
	}

	@keyframes dotScaleOpacity {
		0%, 80%, 100% {
			transform: scale(0);
			opacity: 0;
		}
		40% {
			transform: scale(1);
			opacity: 1;
		}
	}
</style>
