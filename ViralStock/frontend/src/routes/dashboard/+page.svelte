<script lang="ts">
	let keyword = '';
	let loading = false;

	let results: Array<{ thumbnailUrl: string; videoUrl: string }> = [];

	async function search() {
		if (!keyword.trim()) return;

		loading = true;
		results = [];

		try {
			const response = await fetch('/api/search', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ keyword })
			});

			if (!response.ok) {
				console.error('Search failed:', response.statusText);
				// In a real app, you'd show a user-facing error here
			} else {
				results = await response.json();
			}
		} catch (error) {
			console.error('An error occurred:', error);
		} finally {
			loading = false;
		}
	}

	async function downloadVideo(videoUrl: string) {
		try {
			const response = await fetch('/api/download', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ videoUrl })
			});

			if (!response.ok) {
				throw new Error(`Download failed: ${response.statusText}`);
			}

			const blob = await response.blob();
			const url = window.URL.createObjectURL(blob);
			const a = document.createElement('a');
			a.style.display = 'none';
			a.href = url;
			// You could generate a more descriptive filename here
			a.download = 'viralstock-video.mp4';
			document.body.appendChild(a);
			a.click();
			window.URL.revokeObjectURL(url);
			a.removeChild(a);
		} catch (error) {
			console.error('Download error:', error);
			// Optionally, show an error message to the user
		}
	}
</script>

<div class="flex min-h-screen flex-col items-center bg-gray-50 p-4">
	<div class="w-full max-w-2xl">
		<h1 class="mb-4 text-center text-4xl font-bold">ViralStock</h1>
		<div class="flex">
			<input
				type="text"
				bind:value={keyword}
				placeholder="Enter a keyword..."
				class="flex-grow rounded-l-md border border-gray-300 p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
				on:keydown={(e) => e.key === 'Enter' && search()}
			/>
			<button
				on:click={search}
				class="rounded-r-md bg-blue-500 px-4 py-2 text-white hover:bg-blue-600"
				disabled={loading}
			>
				{#if loading}
					<span>...</span>
				{:else}
					Search
				{/if}
			</button>
		</div>

		{#if loading}
			<div class="mt-8 w-full">
				<div class="h-2 w-full overflow-hidden rounded-full bg-gray-200">
					<div class="h-full animate-pulse rounded-full bg-blue-500" style="width: 50%" />
				</div>
				<p class="mt-2 text-center text-sm text-gray-500">Searching for videos...</p>
			</div>
		{/if}

		{#if results.length > 0}
			<div class="mt-8 grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
				{#each results as result (result.videoUrl)}
					<div class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm">
						<img
							src={result.thumbnailUrl}
							alt="Video thumbnail"
							class="h-48 w-full object-cover"
						/>
						<div class="p-4">
							<button
								on:click={() => downloadVideo(result.videoUrl)}
								class="w-full rounded-md bg-green-500 px-4 py-2 text-white hover:bg-green-600"
							>
								Download
							</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
