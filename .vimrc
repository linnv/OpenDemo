set encoding=utf-8
set ff=unix
set ffs=unix,dos
set ruler "display the current row and column on the right_bottom corner 
set relativenumber
set nu

" set clipboard=unnamed
set clipboard=unnamedplus

" yaml indentation
au FileType yaml setlocal tabstop=4 expandtab shiftwidth=4 softtabstop=4

" 空格代替Tab 
" set expandtab
" 缩进宽度
" set tabstop=4 
" set shiftwidth=4
" 禁在Makefile 中将Tab转换成空格
autocmd FileType makefile set noexpandtab
autocmd FileType md set noexpandtab

" set foldmethod=syntax "fold based on indent
set foldnestmax=10      "deepest fold is 10 levels
set nofoldenable        "dont fold by default
set foldlevel=1         "this is just what i use

set fillchars+=vert:│
" set fillchars+=vert:\ 
set foldcolumn=0
hi LineNr ctermbg=white ctermfg=black
" hi LineNr ctermbg=white ctermfg=red
hi VertSplit ctermbg=black ctermfg=white

set hlsearch
" set backspace=indent,eol,start

" ignorecase and smartcase together make Vim deal with case-sensitive search intelligently. If you search for an all-lowercase string your search will be case-insensitive, but if one or more characters is uppercase the search will be case-sensitive. Most of the time this does what you want.
set ignorecase
set smartcase
set showmatch
set showmode
set showcmd
set gdefault

set cursorline
set cursorcolumn

" set colorcolumn=85
set wrap
set mouse=a

set timeoutlen=1000 ttimeoutlen=0

set modelines=0   "don't execute command that comment in a file"
set undofile  "this allow you to undo change even when you reopen a file
set laststatus=2   " Always show the statusline

"remember last edicting position
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif

" ====================plugin manager=================
call plug#begin('~/.vim/plugged')
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'scrooloose/nerdtree'
let NERDTreeIgnore=['node_modules$[[dir]]']
" enable line numbers
let NERDTreeShowLineNumbers=1
" make sure relative line numbers are used
autocmd FileType nerdtree setlocal relativenumber
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code completion> start""""""""""""""""""""""""""""""""""
Plug 'Valloric/YouCompleteMe',{'do': './install.py --clang-completer'}
" Plug 'Valloric/YouCompleteMe',{'do': './install.py --tern-completer'}
" let g:ycm_global_ycm_extra_conf = '/home/jialin/archive/CutDemo/.ycm_extra_conf.py'

let g:ycm_filetype_blacklist = {}
let g:ycm_auto_trigger = 1
let g:ycm_min_num_of_chars_for_completion = 2
let g:ycm_key_list_select_completion = ['<c-j>','<Down>']
let g:ycm_key_list_previous_completion= ['<c-k>','<Up>']
let g:ycm_key_invoke_completion = '<c-z>'
" let g:ycm_python_binary_path = 'python2.7'
let g:ycm_confirm_extra_conf = 0
nmap <M-g> :YcmCompleter GoToDefinitionElseDeclaration <C-R>=expand("<cword>")<CR><CR>  
Plug 'rdnetto/YCM-Generator', { 'branch': 'stable'}
"""""""""""""""""""""""""""""<code completion> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code snippets> start""""""""""""""""""""""""""""""""""
Plug 'SirVer/ultisnips' | Plug 'honza/vim-snippets'
let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger = "<c-l>"
let g:UltiSnipsJumpBackwardTrigger="<c-h>"
let g:UltiSnipsEditSplit="vertical"
"use absolute path 
let g:UltiSnipsSnippetsDir = '~/.vim/UltiSnips'
"""""""""""""""""""""""""""""<code snippets> end""""""""""""""""""""""""""""""""""""

"for js
Plug 'ternjs/tern_for_vim'
Plug 'pangloss/vim-javascript'
Plug 'posva/vim-vue'
Plug 'othree/html5.vim'

"auto format js html
" au BufNewFile,BufRead *.vue set filetype=html
autocmd BufWritePre *.{js,html,htm,vue} :normal migg=G`izz
" autocmd BufWritePre *.{js,html,htm} :normal ggVG=

" Markdown
autocmd BufNewFile,BufRead *.{md,mkd,mkdn,mark*}  nested setlocal filetype=markdown

"for django
au BufNewFile,BufRead *.hj set filetype=htmldjango

Plug 'godlygeek/tabular'
Plug 'plasticboy/vim-markdown'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"neomake
Plug 'neomake/neomake'
autocmd! BufWritePost * Neomake
" autocmd BufWritePost * Neomake
let g:neomake_open_list=2
let g:neomake_go_enabled_makers = [ 'go', 'gometalinter' ]
let g:neomake_go_gometalinter_maker = {
  \ 'args': [
  \   '--tests',
  \   '--enable-gc',
  \   '--concurrency=3',
  \   '--fast',
  \   '-D', 'gocyclo',
  \   '-D', 'gotype',
  \   '-D', 'dupl',
  \   '-D', 'golint',
  \   '-D', 'gas',
  \   '-D', 'ineffassign',
  \   '-E', 'interfacer',
  \   '-E', 'goconst',
  \   '-E', 'aligncheck',
  \   '-E', 'unconvert',
  \   '-E', 'errcheck',
  \   '-E', 'misspell',
  \   '-E', 'unused',
  \   '-D', 'vet',
  \   '-D', 'vetshadow',
  \   '%:p:h',
  \ ],
  \ 'append_file': 0,
  \ 'errorformat':
  \   '%E%f:%l:%c:%trror: %m,' .
  \   '%W%f:%l:%c:%tarning: %m,' .
  \   '%E%f:%l::%trror: %m,' .
  \   '%W%f:%l::%tarning: %m'
  \ }
" vue is treated as html file but tidy does'nt reconize vue syntax, so
" disabled tidy for vue safty
let g:neomake_html_enabled_makers = []
let g:neomake_css_enabled_makers = []

"tags lister method menu
Plug 'majutsushi/tagbar'
" let g:tagbar_width = 30
" let g:tagbar_left = 1
let g:tagbar_show_linenumbers = 2
let g:tagbar_type_markdown = {
        \ 'ctagstype' : 'markdown',
        \ 'kinds' : [
                \ 'h:headings',
        \ ],
    \ 'sort' : 0
\ }
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"shortcut for comment,e.g. gcc
Plug 'tomtom/tcomment_vim'
Plug 'tpope/vim-surround'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"json
" Plug 'elzr/vim-json'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"power search tool
Plug 'mileszs/ack.vim'
if executable('ag')
  let g:ackprg = 'ag --vimgrep'
endif

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"show git diff when edicting
Plug 'airblade/vim-gitgutter'
let g:gitgutter_signs = 1
let g:gitgutter_highlight_lines = 0

"""""""""""""""""""""""""""""<golang tool chain> start""""""""""""""""""""""""""""""""""
" Plug 'fatih/vim-go'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
let g:go_def_mode='gopls'
let g:go_info_mode='gopls'
let g:go_fmt_command = 'goimports'    "auto insert package
" let g:go_gocode_propose_source=0
"horizon
au FileType go nmap <Leader>h <Plug>(go-def-split) 
au FileType go nmap <Leader>v <Plug>(go-def-vertical)
au FileType go nmap <Leader>t <Plug>(go-def-tab)
"""""""""""""""""""""""""""""<golang tool chain> end""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" select multiple words and update them
Plug 'terryma/vim-multiple-cursors'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"scoll smoothly when hit c-b c-f
" Plug 'yonchu/accelerated-smooth-scroll'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"powerline like
Plug 'itchyny/lightline.vim'
let g:lightline = {
      \ 'colorscheme': 'solarized',
      \ 'component_function': {
      \   'filename': 'LightLineFilename'
      \ },
      \ }
function! LightLineFilename()
  " return expand('%:p')
  return expand('%')
endfunction

" set dropdown to match solarized light
highlight Pmenu ctermfg=254 ctermbg=33
highlight PmenuSel ctermfg=235 ctermbg=39 cterm=bold
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"input stat switch between normal mode and insert mode automatically
Plug 'CodeFalling/fcitx-vim-osx'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" visual start search
Plug 'bronson/vim-visual-star-search'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'raimondi/delimitmate'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'ctrlpvim/ctrlp.vim'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'editorconfig/editorconfig-vim'
" let g:EditorConfig_exec_path ='~/.config/nvim/plugged/editorconfig-vim/plugin/editorconfig-core-py'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Plug 'justinmk/vim-sneak'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Plug 'easymotion/vim-easymotion'
" Move to word
" map  w <Plug>(easymotion-bd-w)
" nmap w <Plug>(easymotion-overwin-w)
" s{char}{char} to move to {char}{char}
" nmap s <Plug>(easymotion-overwin-f2)
" Move to line
" map l <Plug>(easymotion-bd-jk)
" nmap l <Plug>(easymotion-overwin-line)

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'buoto/gotests-vim'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'hotoo/pangu.vim'
autocmd BufWritePre *.markdown,*.md,*.text,*.txt,*.wiki,*.cnx call PanGuSpacing()
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'altercation/vim-colors-solarized'
set background=light
" let g:solarized_visibility = "high"
" let g:solarized_contrast = "high"
" let g:solarized_termcolors=256

" colorscheme molokai
" let g:molokai_original = 1
" let g:rehash256 = 1
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'ryanoasis/vim-devicons'
let g:airline_powerline_fonts = 1

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:vimrc_author='Jialin Wu' 
let g:vimrc_email='win27v@gmail.com' 
let g:vimrc_homepage='https://jialinwu.com'

call plug#end()

" ====================leader key=================
 "Set mapleader
let mapleader = "\<Space>"

"show file edicted recently 
" set runtimepath^=~/.config/nvim/bundle/ctrlp.vim
nnoremap <silent> <leader>r :CtrlPMRU<cr>
" map <silent> <leader>t :CtrlPMixed<cr>
" map <silent> <leader>t :CtrlP<cr>

"list file in current directory
Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
" set rtp+=/Users/Jialin/.fzf
nnoremap <silent> <leader>f :FZF<cr>

"Fast opening nerdtree 
nnoremap <silent> <leader>nt :NERDTree<cr>

"Save and exit
map <silent> <leader>x :x<cr>
"select all
map <silent> <leader>a ggvGV<cr>

"exit without saving
map <silent> <leader>q :q!<cr>

map <silent> <leader>b :NERDTreeFromBookmark 

noremap <Leader>d "+d
noremap <Leader>y "+y
noremap <Leader>1y "1y
noremap <Leader>2y "2y

noremap <Leader>p "+p
noremap <Leader>P "+P
noremap <Leader>1p "1p
noremap <Leader>1P "1P

noremap <Leader>2p "2p
noremap <Leader>2P "2P

"write zone
map <silent> <leader>wz :Goyo<cr>

"exit write zone
map <silent> <leader>ewz :Goyo!<cr>

"shwo absolute filename
nnoremap <silent> <leader>af :echo expand('%:p')<cr>

" quick into searching
nnoremap <silent> <leader>k :Ack 

"new tab 
nnoremap <silent> <leader>tn :tabnew<cr>
"close one tab 
nnoremap <silent> <leader>tc :tabclose<cr>

"switch window area
"for terminal of nvim
" :tnoremap <C-h> <C-\><C-n><C-w>h
" :tnoremap <C-j> <C-\><C-n><C-w>j
" :tnoremap <C-k> <C-\><C-n><C-w>k
" :tnoremap <C-l> <C-\><C-n><C-w>l
:tnoremap <C-\> <C-\><C-n>
"don't use this map while c-c is useful in terminal to operate shell cmd
" :tnoremap <C-c> <C-\><C-n>

:nnoremap <C-h> <C-w>h
:nnoremap <C-j> <C-w>j
:nnoremap <C-k> <C-w>k
:nnoremap <C-l> <C-w>l
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

nnoremap ; : 
nmap    w=  :resize +20<CR>
nmap    w-  :resize -20<CR>
nmap    w,  :vertical resize -20<CR>
nmap    w.  :vertical resize +20<CR>
"""""""""""""""""""""""""""""<shortcut defined> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<for cpp/c> start""""""""""""""""""""""""""""""""""
" switch between header/source with  hc for cpp
map <silent> <leader>hc :e %:p:s,.h$,.X123X,:s,.cpp$,.h,:s,.X123X$,.cpp,<cr>
"""""""""""""""""""""""""""""<for cpp/c> end""""""""""""""""""""""""""""""""""""

"edit tmp file
map <silent> <leader>tmp :e ~/tmpfile<cr>
map <silent> <leader>booknote :e /Users/Jialin/Dropbox/vimNote/book.md<cr>
map <silent> <leader>jsnote :e /Users/Jialin/Dropbox/vimNote/jsnote.md<cr>

"hightlight and search
vnoremap // y/<C-R>"<CR>

"make
map <silent> <leader>zm :make <cr>
map <silent> <leader>zl :lopen<cr>
map <silent> <leader>m :TagbarToggle<cr>

"tab to %
nnoremap <tab> %
vnoremap <tab> %
"keep hightlighting while doing indent 
vnoremap < <gv
vnoremap > >gv

"selec text from current cursor to the end of line
nnoremap Y v$

"delete text from current cursor to the end or begining of line in insert mode
map! <M-l>  <ESC>v$c
map! <M-h>  <ESC>v^c

"Keep search pattern at the center of the screen
nnoremap <silent> n nzz
nnoremap <silent> N Nzz
nnoremap <silent> * *zz
nnoremap <silent> # #zz
nnoremap <silent> g* g*zz

"Better comand-line editing
cnoremap <C-a> <Home>
cnoremap <C-e> <End>

"replace : with ; in normal mode
nnoremap ; :
"f{xx} forwrad`,` backward`<`
nnoremap ' ;
nnoremap " ,

"split window
nnoremap <leader>2w <C-w>v<C-w>l

"visul for html tab
nnoremap <leader>zv Vat

"use alt-w to save
nnoremap <M-w> :w<cr>
inoremap <M-w> <ESC>:w<cr>

"use jj to exit back to normal mode
inoremap lk <ESC>

"Save file edicting
nnoremap <silent> <leader>w :w<cr>

"Fast editing of edictor configuration
map <silent> <leader>ee :e ~/.vimrc<cr>
autocmd! bufwritepost .vim.rc source ~/.vimrc

hi Directory ctermfg=Blue
" hi Directory ctermfg=lightBlue
" hi Directory ctermfg=grey

" let g:python_host_prog='/usr/local/opt/python2/libexec/bin/python'
" let g:python3_host_prog='/usr/local/opt/python3/libexec/bin/python'
let g:python_host_skip_check = 1

