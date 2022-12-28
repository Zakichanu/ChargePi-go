<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import {
		faHome,
		faBolt,
		faCreditCard,
		faDisplay,
		faCloud,
		faFilePen,
		faRightFromBracket,
		faBars,
	} from '@fortawesome/free-solid-svg-icons';
	import NavItem from '$lib/navigation/NavItem.svelte';

	let expand = true;
</script>

<div class="flex flex-col h-full space-y-4 p-4">
	<div class="navbar rounded-box bg-base-200 shadow-sm">
		<div class="flex-1">
			<span class="btn btn-ghost normal-case text-xl">ChargePI</span>
		</div>
		<div class="flex-none">
			<button class="btn btn-square btn-ghost" on:click={() => (expand = !expand)}>
				<Fa icon={faBars} />
			</button>
		</div>
	</div>

	<div class="flex h-full space-x-4">
		<div class="card bg-base-200 shadow-sm grow p-4">
			<slot />
		</div>
		{#if expand}
			<div class="card bg-base-200 shadow-sm">
				<ul class="menu w-64 p-4">
					<li>
						<NavItem href="/">
							<Fa icon={faHome} />
							Dashboard
						</NavItem>
					</li>
					<li>
						<NavItem href="/evse">
							<Fa icon={faBolt} />
							EVSE
						</NavItem>
					</li>
					<li>
						<NavItem href="/card-reader">
							<Fa icon={faCreditCard} />
							Card Reader
						</NavItem>
					</li>
					<li>
						<NavItem href="/display">
							<Fa icon={faDisplay} />
							Display
						</NavItem>
					</li>
					<li>
						<NavItem href="/ocpp">
							<Fa icon={faCloud} />
							OCPP
						</NavItem>
					</li>
					<li>
						<NavItem href="/logs">
							<Fa icon={faFilePen} />
							Logs
						</NavItem>
					</li>
					<li>
						<NavItem href="/login">
							<Fa icon={faRightFromBracket} />
							Logout
						</NavItem>
					</li>
				</ul>
			</div>
		{/if}
	</div>
</div>
