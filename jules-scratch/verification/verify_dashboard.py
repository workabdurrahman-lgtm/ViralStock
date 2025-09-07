from playwright.sync_api import sync_playwright, expect

def run():
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        page = browser.new_page()
        page.goto("http://localhost:5173/dashboard")

        # Expect the main heading to be visible to ensure the page has loaded
        expect(page.get_by_role("heading", name="ViralStock")).to_be_visible()

        page.screenshot(path="jules-scratch/verification/verification.png")
        browser.close()

run()
