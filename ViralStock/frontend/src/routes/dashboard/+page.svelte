<script lang="ts">
  import { user } from '$lib/stores/auth';
  import { auth } from '$lib/firebase';
  import { signOut } from 'firebase/auth';
  import { goto } from '$app/navigation';

	let keyword = '';
	let loading = false;

	let results: Array<{ thumbnailUrl: string; videoUrl: string }> = [];

	async function search() {
		if (!keyword.trim()) return;

		loading = true;
		results = [];

		try {
			const response = await fetch('http://localhost:8080/api/search', {
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
			const response = await fetch('http://localhost:8080/api/download', {
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

{#if $user === undefined}
  <div class="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4">
    <p class="text-lg">Loading...</p>
  </div>
{:else if $user}
  <div class="flex min-h-screen flex-col items-center bg-gray-50 p-4">
    <div class="w-full max-w-2xl">
      <div class="flex justify-between items-center mb-4">
        <h1 class="text-center text-4xl font-bold">ViralStock</h1>
        <button
          on:click={() => signOut(auth)}
          class="rounded-md bg-gray-200 px-4 py-2 text-gray-800 hover:bg-gray-300"
        >
          Logout
        </button>
      </div>
      <div class="flex">
        <input
          type="text"
          bind:value={keyword}
          placeholder="Enter a keyword..."
          class="flex-grow rounded-l-md border-2 border-green-500 p-2 focus:outline-none focus:ring-2 focus:ring-green-500"
          on:keydown={(e) => e.key === 'Enter' && search()}
        />
        <button
          on:click={search}
          class="rounded-r-md bg-green-500 px-4 py-2 text-white hover:bg-green-600 border-2 border-green-500"
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
            <div class="h-full animate-pulse rounded-full bg-green-500" style="width: 50%"></div>
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
                  class="w-full rounded-md bg-green-500 px-4 py-2 text-white hover:bg-green-600 border-2 border-green-500"
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
{:else}
  <div class="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4">
    <div class="w-full max-w-md text-center">
      <h1 class="mb-4 text-2xl font-bold">Access Denied</h1>
      <p class="mb-8 text-gray-600">You must be logged in to use this feature.</p>
      <button
        on:click={() => goto('/login')}
        class="rounded-md bg-green-500 px-6 py-3 text-white font-semibold shadow-md transition-colors hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
      >
        Login
      </button>
    </div>
  </div>
{/if}
