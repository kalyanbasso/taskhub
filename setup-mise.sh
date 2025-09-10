#!/bin/bash

# Setup script for TaskHub with mise

set -e

echo "🚀 Setting up TaskHub with mise..."

# Check if mise is installed
if ! command -v mise &> /dev/null; then
    echo "❌ mise is not installed. Installing..."
    curl https://mise.run | sh
    echo "✅ mise installed successfully"
    
    # Add mise to PATH for current session
    export PATH="$HOME/.local/bin:$PATH"
    
    echo "📝 Please restart your terminal or run: source ~/.bashrc (or ~/.zshrc)"
    echo "   Then run this script again."
    exit 0
fi

# Check if we're in the project directory
if [ ! -f ".mise.toml" ]; then
    echo "❌ Please run this script from the TaskHub project directory"
    exit 1
fi

echo "🔐 Trusting mise configuration..."
mise trust

echo "📦 Installing tools (Go 1.23)..."
mise install

echo "🔍 Checking installation..."
mise current

echo ""
echo "✅ Setup complete! You can now use:"
echo "   mise tasks          - List available tasks"
echo "   mise run docker-up  - Start Docker containers"
echo "   mise run run        - Run the application locally"
echo "   mise run test       - Run tests"
echo ""
echo "📖 For more information, see MISE.md"
