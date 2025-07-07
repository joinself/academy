(function() {
    'use strict';
    
    // Process GitHub embeds with direct script injection
    function processGitHubEmbeds() {
        const embeds = document.querySelectorAll('[data-github-embed]:not([data-processed])');
        
        if (embeds.length > 0) {
            console.log(`🔍 GitHub Embeds: Found ${embeds.length} new embeds to process`);
        }
        
        embeds.forEach((embed, index) => {
            const target = embed.getAttribute('data-github-embed');
            if (!target) {
                console.warn('⚠️ GitHub embed missing data-github-embed attribute');
                return;
            }
            
            console.log(`📄 Processing embed: ${target}`);
            
            const style = embed.getAttribute('data-style') || 'github-dark';
            const showBorder = embed.getAttribute('data-show-border') !== 'false';
            const showLineNumbers = embed.getAttribute('data-show-line-numbers') !== 'false';
            const showFileMeta = embed.getAttribute('data-show-file-meta') !== 'false';
            const showFullPath = embed.getAttribute('data-show-full-path') !== 'false';
            const showCopy = embed.getAttribute('data-show-copy') !== 'false';
            
            try {
                // Mark as processed immediately to avoid duplicates
                embed.setAttribute('data-processed', 'true');
                
                // Build the parameters string manually to avoid encoding issues
                let paramStr = `target=${encodeURIComponent(target)}`;
                paramStr += `&style=${style}`;
                paramStr += `&type=code`;
                paramStr += `&showBorder=${showBorder ? 'on' : 'off'}`;
                paramStr += `&showLineNumbers=${showLineNumbers ? 'on' : 'off'}`;
                paramStr += `&showFileMeta=${showFileMeta ? 'on' : 'off'}`;
                paramStr += `&showFullPath=${showFullPath ? 'on' : 'off'}`;
                paramStr += `&showCopy=${showCopy ? 'on' : 'off'}`;
                
                const scriptUrl = `https://emgithub.com/embed-v2.js?${paramStr}`;
                console.log(`🔗 Script URL: ${scriptUrl}`);
                
                // Clear the loading message and create script element
                embed.innerHTML = '';
                const script = document.createElement('script');
                script.src = scriptUrl;
                script.async = true;
                
                script.onload = function() {
                    console.log(`✅ GitHub embed loaded: ${target}`);
                };
                
                script.onerror = function(error) {
                    console.error(`❌ Failed to load GitHub embed: ${target}`, error);
                    embed.innerHTML = `<div style="padding: 1rem; border: 1px solid #444; background: #1e1e1e; color: #fff; border-radius: 4px;">
                        <p>⚠️ Failed to load code example</p>
                        <p><a href="${target}" target="_blank" style="color: #58a6ff;">View on GitHub →</a></p>
                    </div>`;
                };
                
                // Append script to the embed container
                embed.appendChild(script);
                
            } catch (error) {
                console.error('❌ Error processing GitHub embed:', error, 'Target:', target);
                embed.innerHTML = `<div style="padding: 1rem; border: 1px solid #444; background: #1e1e1e; color: #fff; border-radius: 4px;">
                    <p>⚠️ Error loading code example</p>
                    <p><a href="${target}" target="_blank" style="color: #58a6ff;">View on GitHub →</a></p>
                </div>`;
                embed.setAttribute('data-processed', 'true');
            }
        });
    }
    
    // Initialize when DOM is ready
    function init() {
        console.log('🚀 Initializing GitHub embeds...');
        processGitHubEmbeds();
    }
    
    // Run on initial load
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
    
    // Handle instant navigation for Material MkDocs
    if (window.location.hash) {
        setTimeout(processGitHubEmbeds, 100);
    }
    
    // Handle tab switching for Material MkDocs
    function handleTabSwitching() {
        const tabButtons = document.querySelectorAll('.tabbed-set input[type="radio"]');
        tabButtons.forEach(button => {
            button.addEventListener('change', function() {
                // Process embeds when tab becomes active
                setTimeout(processGitHubEmbeds, 50);
            });
        });
    }
    
    // Fallback for instant navigation and dynamic content
    const observer = new MutationObserver((mutations) => {
        let shouldProcess = false;
        mutations.forEach(mutation => {
            if (mutation.type === 'childList') {
                // Check if new nodes contain github embeds or tab content
                mutation.addedNodes.forEach(node => {
                    if (node.nodeType === 1) { // Element node
                        if (node.hasAttribute && node.hasAttribute('data-github-embed')) {
                            shouldProcess = true;
                        } else if (node.querySelector && node.querySelector('[data-github-embed]')) {
                            shouldProcess = true;
                        } else if (node.classList && (node.classList.contains('tabbed-set') || node.classList.contains('tabbed-content'))) {
                            shouldProcess = true;
                        }
                    }
                });
            }
        });
        
        if (shouldProcess) {
            setTimeout(processGitHubEmbeds, 100);
        }
    });
    
    observer.observe(document.body, {
        childList: true,
        subtree: true
    });
    
    // Also handle tab switching when DOM is ready
    setTimeout(handleTabSwitching, 500);
})(); 
