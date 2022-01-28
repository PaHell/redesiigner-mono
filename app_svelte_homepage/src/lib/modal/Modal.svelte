<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import Btn from '$lib/Button.svelte';
	import { createEventDispatcher } from 'svelte';
</script>

<script lang="ts">
	export let icon = 'home';
	export let title = 'Title';
	export let text = 'Text';
	export let color = 'accent';
	export let buttons: TButton<(close: () => void) => void, void>[] = [];
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
	<menu
		class="app-modal app-modal-{color} animate__animated {visible
			? 'animate__bounceIn'
			: 'animate__bounceOut'}"
	>
		<main>
			<Icon name={icon} large circle />
			<div class="mt-3 sm:mt-0 sm:ml-3 text-center sm:text-left">
				<h3
					class="mt-4 sm:mt-0 text-xl sm:text-lg leading-6 font-semibold text-gray-900 dark:text-white"
				>
					{title}
				</h3>
				<div class="mt-2 sm:mt-1 flex justify-center">
					<p class="text-base sm:text-sm text-gray-500 dark:text-gray-400 max-w-sm sm:max-w-none">
						{text}
					</p>
				</div>
				<slot />
			</div>
		</main>
		<footer class="app-bar-button border-t dark:border-t-2 dark:border-gray-700">
			{#each buttons as button}
				<Btn
					text={button.text}
					icon={button.icon}
					color={button.color}
					on:click={(e) => {
						button.OnClick(close);
					}}
				/>
			{/each}
		</footer>
	</menu>
</template>

<style global lang="postcss">
	.app-modal {
		@apply rounded-xl sm:rounded-lg overflow-hidden
        bg-gray-50 border border-gray-300 shadow transform transition-all;

		& > main {
			@apply sm:flex sm:items-start
            p-4 sm:p-6 py-6 sm:pr-5;

			& > .app-icon {
				@apply mx-auto sm:mx-0 text-gray-500 bg-gray-100;
			}
		}

		/* default */
		& > main > .app-icon {
			@apply text-gray-500 bg-gray-200;
		}

		/* colors */
		&.app-modal-accent {
			& > main > .app-icon {
				@apply bg-accent-100 text-accent-500;
			}
		}
		&.app-modal-success {
			& > main > .app-icon {
				@apply bg-success-100 text-success-500;
			}
		}
		&.app-modal-warning {
			& > main > .app-icon {
				@apply bg-warning-100 text-warning-500;
			}
		}
		&.app-modal-danger {
			& > main > .app-icon {
				@apply bg-danger-100 text-danger-500;
			}
		}
		&.app-modal-info {
			& > main > .app-icon {
				@apply bg-info-100 text-info-500;
			}
		}
	}

	html.dark .app-modal {
		@apply bg-gray-800 shadow-none border-gray-700 border-2 ring-2 ring-gray-900;

		/* default */
		& > main > .app-icon {
			@apply text-gray-400 bg-transparent;
		}

		/* colors */
		&.app-modal-accent {
			& > main > .app-icon {
				@apply text-accent-400;
			}
		}
		&.app-modal-success {
			& > main > .app-icon {
				@apply text-success-400;
			}
		}
		&.app-modal-warning {
			& > main > .app-icon {
				@apply text-warning-400;
			}
		}
		&.app-modal-danger {
			& > main > .app-icon {
				@apply text-danger-400;
			}
		}
		&.app-modal-info {
			& > main > .app-icon {
				@apply text-info-400;
			}
		}
	}
</style>
