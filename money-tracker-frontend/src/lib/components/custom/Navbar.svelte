<script lang="ts">
	import AvatarFallback from '../ui/avatar/avatar-fallback.svelte';
	import AvatarImage from '../ui/avatar/avatar-image.svelte';
	import Avatar from '../ui/avatar/avatar.svelte';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import ModalForm from './ModalForm.svelte';

	type props = {
		name?: string;
		status?: boolean;
		auth?: boolean;
	};

	let openModal = $state(false);
	let { name = 'joko', status = false, auth = false }: props = $props();
</script>

<div class="border-8 border-black bg-white px-6 py-4">
	<div class="grid grid-cols-2">
		<div class="ml-4 flex items-center justify-start gap-6">
			<div
				class="-rotate-1 transform border-4 border-black bg-black px-4 py-2 shadow-[4px_4px_0px_0px_#666666]"
			>
				<p class="text-base font-black tracking-wider text-white uppercase">MONEY TRACKER</p>
			</div>
			<div class="flex flex-1 flex-row space-x-2">
				<Input
					type="text"
					placeholder="SEARCH YOUR EXPENSES..."
					class="h-10 flex-1 rotate-1 transform border-4 border-black bg-white font-bold tracking-wide text-black uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 placeholder:text-gray-500 focus:rotate-0 focus:border-gray-800 focus:ring-0"
				/>
				<Button
					variant="outline"
					class="-rotate-1 transform border-4 border-black bg-gray-800 font-black tracking-wider text-white uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-black hover:shadow-[2px_2px_0px_0px_#333333]"
				>
					🔍 FIND
				</Button>
			</div>
		</div>
		<div class="mr-4 flex items-center justify-end gap-3">
			{#if auth}
				<div
					class="rotate-1 transform border-4 border-black bg-white p-2 shadow-[6px_6px_0px_0px_#666666]"
				>
					<Avatar class="h-10 w-10 border-2 border-black transition-all duration-200">
						<AvatarImage src="icon.jpg" alt={`${name} avatar`} />
						<AvatarFallback class="border-2 border-black bg-gray-800 font-black text-white">
							{name?.charAt(0).toUpperCase()}
						</AvatarFallback>
					</Avatar>
				</div>
				<div
					class="flex min-w-0 -rotate-1 transform flex-col justify-center border-4 border-black bg-gray-200 p-3 shadow-[6px_6px_0px_0px_#666666]"
				>
					<p class="truncate text-sm leading-tight font-black tracking-wide text-black uppercase">
						{name}
					</p>
					<p
						class="flex items-center gap-2 text-xs leading-tight font-bold tracking-wide uppercase"
					>
						<span
							class="h-3 w-3 border-2 border-black {status === true
								? 'bg-white'
								: status === false
									? 'bg-gray-800'
									: 'bg-gray-400'}"
						></span>
						<span class="text-black">
							{status ? 'ACTIVE' : 'INACTIVE'}
						</span>
					</p>
				</div>
			{:else}
				<Button
					size="sm"
					variant="outline"
					class="rotate-1 transform border-6 border-black bg-white font-black tracking-wider text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:bg-gray-800 hover:text-white hover:shadow-[3px_3px_0px_0px_#333333]"
					onclick={() => {
						openModal = !openModal;
					}}
				>
					🚀 SIGN IN
				</Button>
			{/if}
		</div>
	</div>
</div>

<ModalForm bind:open={openModal} />
