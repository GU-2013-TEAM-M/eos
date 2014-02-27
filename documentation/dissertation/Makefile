

TeamM_Dissertation.pdf: TeamM_Dissertation.tex TeamM_Dissertation.bib
	pdflatex TeamM_Dissertation
	bibtex TeamM_Dissertation
	pdflatex TeamM_Dissertation
	pdflatex TeamM_Dissertation

clean:
	rm -f *.log *.aux *.toc *.out
	rm -f *.bbl *.blg
	rm -f TeamM_Dissertation.pdf
	
