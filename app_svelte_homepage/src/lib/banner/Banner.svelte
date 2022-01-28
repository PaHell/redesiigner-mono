<script lang="ts" context="module">
	import Btn from '../Button.svelte';
	import Icon from '../Icon.svelte';
	import { createEventDispatcher } from 'svelte';
</script>

<script lang="ts">
	export let icon = 'home';
	export let text = 'Lorem ipsum dolor Sit amet.';
	export let color = 'default';
	export let buttons: TButton<void, void>[] = [];
	export let remove = () => {};

	let visible = true;
	function close(): void {
		visible = false;
		setTimeout(() => {
			visible = true;
			remove();
		}, 500);
	}
</script>

<template>
	<div
		class="app-banner app-banner-{color} animate__animated {visible
			? 'animate__fadeInDown'
			: 'animate__fadeOutUp'}"
		class:app-visible={visible}
	>
		<Icon name={icon} />
		<p class="font-medium">{text}</p>
		<div class="app-bar-button bg-transparent">
			{#each buttons as button}
				<Btn
					text={button.text}
					icon={button.icon}
					color={!color || color === 'default' ? 'default' : button.color}
					on:click={(e) => {
						button.OnClick(close);
					}}
				/>
			{/each}
			<Btn icon="cross" color="transparent-default" on:click={(e) => close()} />
		</div>
	</div>
</template>

<style global lang="postcss">
	.app-banner {
		@apply w-full max-h-0
    flex justify-center items-center
    transition-all overflow-hidden;

		&.app-visible {
			@apply max-h-20 border-y;
		}

		&:last-child {
			@apply border-t-0;
		}

		& > *:not(:last-child) {
			@apply mr-4;
		}

		& > .app-icon {
			@apply ml-4 p-2 rounded-lg;
		}

		& > p {
			@apply flex-1 text-white;
		}

		/* default */
		&.app-banner-default,
		&.app-banner-undefined {
			@apply bg-gray-100 border-gray-300;
			& > .app-icon {
				@apply bg-gray-200 text-gray-500;
			}
			& > p {
				@apply text-gray-800;
			}
		}
		/* colors */
		&.app-banner-accent {
			@apply bg-accent-600 border-accent-700;
			& > .app-icon {
				@apply bg-accent-800 text-accent-100;
			}
		}
		&.app-banner-success {
			@apply bg-success-600 border-success-700;
			& > .app-icon {
				@apply bg-success-800 text-success-100;
			}
		}
		&.app-banner-warning {
			@apply bg-warning-600 border-warning-700;
			& > .app-icon {
				@apply bg-warning-800 text-warning-100;
			}
		}
		&.app-banner-danger {
			@apply bg-danger-600 border-danger-700;
			& > .app-icon {
				@apply bg-danger-800 text-danger-100;
			}
		}
		&.app-banner-info {
			@apply bg-info-600 border-info-700;
			& > .app-icon {
				@apply bg-info-800 text-info-100;
			}
		}
	}

	html.dark .app-banner {
		& > p {
			@apply text-white;
		}

		/* default */
		&.app-banner-default,
		&.app-banner-undefined {
			@apply bg-gray-800 border-gray-700 border-0;
			& > .app-icon {
				@apply bg-gray-700 text-gray-300;
			}
			&:first-child {
				@apply border-b-2;
			}
		}
		/* colors */
		&.app-banner-accent {
			@apply bg-accent-700 border-accent-800;
			& > .app-icon {
				@apply bg-accent-800;
			}
		}
		&.app-banner-success {
			@apply bg-success-700 border-success-800;
			& > .app-icon {
				@apply bg-success-800;
			}
		}
		&.app-banner-warning {
			@apply bg-warning-700 border-warning-800;
			& > .app-icon {
				@apply bg-warning-800;
			}
		}
		&.app-banner-danger {
			@apply bg-danger-700 border-danger-800;
			& > .app-icon {
				@apply bg-danger-800;
			}
		}
		&.app-banner-info {
			@apply bg-info-700 border-info-800;
			& > .app-icon {
				@apply bg-info-800;
			}
		}
	}
</style>
